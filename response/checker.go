package response

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PaesslerAG/gval"

	"github.com/asalih/guardian/data"
	"github.com/asalih/guardian/models"
)

//Checker Response checker
type Checker struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Response       *http.Response
	Target         *models.Target

	requestTransfer *responseTransfer
	result          chan *models.MatchResult
	firewallResult  chan *models.FirewallMatchResult
	time            time.Time
}

type responseTransfer struct {
	BodyBuffer  []byte
	ContentType string
}

/*NewResponseChecker Request checker initializer*/
func NewResponseChecker(w http.ResponseWriter, r *http.Request, resp *http.Response, target *models.Target) *Checker {
	return &Checker{w, r, resp, target, nil, nil, nil, time.Now()}
}

/*Handle Request checker handler func*/
func (r Checker) Handle() bool {

	if !r.Target.WAFEnabled {
		return false
	}

	//TODO: Open it when checking response body
	//bodyBuffer, _ := ioutil.ReadAll(r.Response.Body)
	//r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBuffer))
	//r.requestTransfer = &requestTransfer{bodyBuffer, r.Request.Header.Get("Content-Type")}

	result := r.handleWAFChecker()

	if result {
		return result
	}

	return r.handleFirewallRuleChecker()
}

func (r Checker) handleFirewallRuleChecker() bool {
	firewallChannel := make(chan bool, 1)
	db := &data.DBHelper{}

	go func() {
		var wg sync.WaitGroup

		firewallRules := db.GetResponseFirewallRules(r.Target.ID)
		lenOfRules := len(firewallRules)

		r.firewallResult = make(chan *models.FirewallMatchResult, lenOfRules)

		wg.Add(lenOfRules)

		mapForFwRules := map[string]interface{}{
			"ip": map[string]interface{}{
				"src": r.Request.RemoteAddr,
			},
			"http": map[string]interface{}{
				"query":    r.Request.URL.RawQuery,
				"path":     r.Request.URL.Path,
				"host":     r.Request.URL.Host,
				"cookie":   models.CookiesToString(r.Request.Cookies()),
				"header":   models.HeadersToString(r.Request.Header),
				"method":   r.Request.Method,
				"protocol": r.Request.Proto,
			},
			"response": map[string]interface{}{
				"status":        r.Response.Status,
				"statusCode":    r.Response.StatusCode,
				"cookie":        models.CookiesToString(r.Response.Cookies()),
				"header":        models.HeadersToString(r.Response.Header),
				"contentLength": r.Response.ContentLength,
			},
		}

		for _, rule := range firewallRules {
			go r.handleFirewallPayload(rule, mapForFwRules, &wg)
		}

		wg.Wait()

		close(r.firewallResult)

		firewallChannel <- true
	}()

	select {
	case <-firewallChannel:
	case <-time.After(3 * time.Minute):
		panic("failed to execute rules.")
	}

	for i := range r.firewallResult {
		//Action: 0 is block
		//Action: 1 is allow
		if i.IsMatched && i.FirewallRule.Action == 0 ||
			!i.IsMatched && i.FirewallRule.Action == 1 {
			r.ResponseWriter.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", r.Request.URL.Path)

			db.LogFirewallMatchResult(i, r.Target, r.Request.RequestURI, true)

			return true
		}
	}

	return false
}

func (r Checker) handleFirewallPayload(rule *models.FirewallRule, mapForFwRules map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	evalResult, everr := gval.Evaluate(rule.Expression, mapForFwRules)

	if everr != nil {
		fmt.Println(everr)
	}

	r.firewallResult <- models.NewFirewallMatchResult(rule, evalResult.(bool)).Time(r.time)
}

func (r Checker) handleWAFChecker() bool {

	done := make(chan bool, 1)

	go func() {
		var wg sync.WaitGroup

		r.result = make(chan *models.MatchResult, models.LenOfGroupedResponsePayloadDataCollection)

		wg.Add(models.LenOfGroupedResponsePayloadDataCollection)

		for key, payload := range models.ResponseCheckPointPayloadData {
			go r.handlePayload(key, payload, &wg)
		}

		wg.Wait()

		close(r.result)

		done <- true
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Minute):
		panic("failed to execute rules.")
	}

	for i := range r.result {
		for _, m := range i.MatchedPayloads {
			if i.IsMatched {
				db := &data.DBHelper{}
				if m.Action == "block" {
					r.ResponseWriter.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", r.Request.URL.Path)

					go db.LogMatchResult(i, m, r.Target, r.Request.RequestURI, true)

					return true
				} else if m.Action == "remove" {
					//Probably action took previously. Just log it
					go db.LogMatchResult(i, m, r.Target, r.Request.RequestURI, true)
				}
				//TODO: Handle new action types
			}
		}
	}

	return false
}

func (r Checker) handlePayload(key string, payloads []models.PayloadData, wg *sync.WaitGroup) {
	defer wg.Done()

	switch key {
	case "Header":
		r.result <- r.handleHeader(payloads)
	}
}

func (r Checker) handleHeader(payloads []models.PayloadData) *models.MatchResult {

	matchResult := models.NewMatchResult(false)
	for _, p := range payloads {
		for hk, hv := range r.Response.Header {
			isMatch, _ := models.IsMatch(p.Payload, hk+": "+hv[0])

			if isMatch {
				if p.Action == "block" {
					matchResult.Append(&p).SetMatch(true).Time(r.time)
					return matchResult
				} else if p.Action == "remove" {
					r.Response.Header.Del(hk)
					matchResult.Append(&p).SetMatch(true).Time(r.time)
				}
			}
		}
	}

	return matchResult
}
