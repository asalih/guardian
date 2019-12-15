package request

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/PaesslerAG/gval"

	"github.com/asalih/guardian/data"
	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/variables"
)

var staticSuffix = []string{".js", ".css", ".png", ".jpg", ".gif", ".bmp", ".svg", ".ico"}

/*Checker Cheks the requests init*/
type Checker struct {
	ResponseWriter  http.ResponseWriter
	Request         *http.Request
	Target          *models.Target
	RequestTransfer *RequestTransfer

	result         []*models.MatchResult
	firewallResult chan *models.FirewallMatchResult
	time           time.Time
}

/*NewRequestChecker Request checker initializer*/
func NewRequestChecker(w http.ResponseWriter, r *http.Request, target *models.Target) *Checker {
	return &Checker{w, r, target, nil, nil, nil, time.Now()}
}

/*Handle Request checker handler func*/
func (r *Checker) Handle() bool {

	if !r.Target.WAFEnabled || r.Request.Method == "GET" && r.IsStaticResource(r.Request.URL.Path) {
		return false
	}

	r.RequestTransfer = NewRequestTransfer(r.Request)
	result := r.handleWAFChecker()

	if result {
		return result
	}

	//return result
	return r.handleFirewallRuleChecker()
}

// IsStaticResource ...
func (r *Checker) IsStaticResource(url string) bool {
	if strings.Contains(url, "?") {
		return false
	}
	for _, suffix := range staticSuffix {
		if strings.HasSuffix(url, suffix) {
			return true
		}
	}
	return false
}

func (r *Checker) handleFirewallRuleChecker() bool {
	firewallChannel := make(chan bool, 1)
	db := &data.DBHelper{}

	go func() {
		var wg sync.WaitGroup

		firewallRules := db.GetRequestFirewallRules(r.Target.ID)
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

			db.LogFirewallMatchResult(i, r.Target, r.Request.RequestURI, false)

			return true
		}
	}

	return false
}

func (r *Checker) handleWAFChecker() bool {

	done := make(chan bool, 1)

	go func() {

		for key, payload := range models.RequestCheckPointPayloadData {
			r.handlePayload(key, payload)
		}

		done <- true
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Minute):
		panic("failed to execute rules.")
	}

	for _, i := range r.result {
		for _, m := range i.MatchedPayloads {
			if i.IsMatched {
				if m.Action == "block" {
					r.ResponseWriter.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", r.Request.URL.Path)

					db := &data.DBHelper{}

					go db.LogMatchResult(i, m, r.Target, r.Request.RequestURI, false)

					return true
				}
				//TODO: Handle new action types if needed
			}
		}

	}

	return false
}

func (r *Checker) handleFirewallPayload(rule *models.FirewallRule, mapForFwRules map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	evalResult, everr := gval.Evaluate(rule.Expression, mapForFwRules)

	if everr != nil {
		fmt.Println(everr)
	}

	r.firewallResult <- models.NewFirewallMatchResult(rule, evalResult.(bool)).Time(r.time)
}

func (r *Checker) handlePayload(key string, payloads []models.PayloadData) {
	switch key {
	case "Query":
		r.result = append(r.result, r.handleQuery(payloads))
	case "Path":
		r.result = append(r.result, r.handlePath(payloads))
	case "Form":
		r.result = append(r.result, r.handleForm(payloads))
	case "Upload":
		//Upload check point will be handled in handleForm function

		r.result = append(r.result, models.NewMatchResult(false))
	default:
		r.result = append(r.result, models.NewMatchResult(false))
	}
}

func (r *Checker) handleQuery(payloads []models.PayloadData) *models.MatchResult {
	matchResult := models.NewMatchResult(false)

	for _, p := range payloads {
		isMatch, _ := models.IsMatch(p.Payload, r.RequestTransfer.Get(variables.QUERY_STRING))

		if isMatch {
			return matchResult.SetMatch(true).Append(&p).Time(r.time)
		}
	}

	return matchResult
}

func (r *Checker) handlePath(payloads []models.PayloadData) *models.MatchResult {
	matchResult := models.NewMatchResult(false)

	for _, p := range payloads {
		isMatch, _ := models.IsMatch(p.Payload, r.RequestTransfer.Get(variables.REQUEST_BASENAME))

		if isMatch {
			return matchResult.SetMatch(true).Append(&p).Time(r.time)
		}
	}

	return matchResult
}

func (r *Checker) handleForm(payloads []models.PayloadData) *models.MatchResult {
	matchResult := models.NewMatchResult(false)

	if strings.HasPrefix(r.RequestTransfer.Get(variables.REQUEST_BODY_TYPE), "multipart/form-data") {
		// ChkPoint_UploadFileExt
		r.Request.ParseMultipartForm(1024)
		if r.Request.MultipartForm != nil {
			for _, filesHeader := range r.Request.MultipartForm.File {
				for _, fileHeader := range filesHeader {
					fileExtension := filepath.Ext(fileHeader.Filename) // .php
					uploadCheck := models.Filter(models.RequestPayloadDataCollection, func(p models.PayloadData) bool { return p.CheckPoint == "Upload" })

					for _, p := range uploadCheck {
						matched, _ := models.IsMatch(p.Payload, fileExtension)
						if matched == true {
							return matchResult.SetMatch(true).Append(&p).Time(r.time)
						}
					}
				}
			}

			// Multipart Content
			multiReader := multipart.NewReader(r.Request.Body, r.RequestTransfer.Get(variables.MULTIPART_BOUNDARY))
			for {
				p, err := multiReader.NextPart()
				if err == io.EOF {
					break
				}
				partContent, err := ioutil.ReadAll(p)

				for _, p := range payloads {
					isMatch, _ := models.IsMatch(p.Payload, models.UnEscapeRawValue(string(partContent)))

					if isMatch {
						return matchResult.SetMatch(true).Append(&p).Time(r.time)
					}
				}
			}
		}

	} else if strings.HasPrefix(r.RequestTransfer.Get(variables.REQUEST_BODY_TYPE), "application/json") {
		var params interface{}

		decoder := json.NewDecoder(r.Request.Body)
		err := decoder.Decode(&params)

		if err != nil {
			panic(err)
		}

		result := r.isJSONValueHitPolicy(payloads, params)
		if result.IsMatched {
			return result
		}
	} else {
		r.Request.ParseForm()
	}

	params := r.Request.Form

	for key, values := range params {

		for _, p := range payloads {
			isMatch, _ := models.IsMatch(p.Payload, key)

			if isMatch {
				return matchResult.SetMatch(true).Append(&p).Time(r.time)
			}

			for _, value := range values {
				if isDigit, err := models.IsMatch(`^\d{1,5}$`, value); err == nil {
					if isDigit {
						continue
					}
				}
				// ChkPoint_ValueLength
				// TODO: Check length
				// valueLength := strconv.Itoa(len(value))
				// matched, policy = IsMatchGroupPolicy(ctxMap, appID, valueLength, models.ChkPointValueLength, "", false)
				// if matched == true {
				// 	return matched, policy
				// }
				// ChkPoint_GetPostValue

				isMatch, _ := models.IsMatch(p.Payload, value)

				if isMatch {
					return matchResult.SetMatch(true).Append(&p).Time(r.time)
				}
			}
		}
	}

	return matchResult
}

func (r *Checker) isJSONValueHitPolicy(payloads []models.PayloadData, value interface{}) *models.MatchResult {
	matchResult := models.NewMatchResult(false)

	valueKind := reflect.TypeOf(value).Kind()
	switch valueKind {
	case reflect.String:
		value2 := value.(string)

		for _, p := range payloads {
			isMatch, _ := models.IsMatch(p.Payload, models.UnEscapeRawValue(value2))

			if isMatch {
				return matchResult.SetMatch(true).Append(&p).Time(r.time)
			}
		}
	case reflect.Map:
		value2 := value.(map[string]interface{})
		for _, subValue := range value2 {
			matched := r.isJSONValueHitPolicy(payloads, subValue)
			if matched.IsMatched {
				return matched
			}
		}
	case reflect.Slice:
		value2 := value.([]interface{})
		for _, subValue := range value2 {
			result := r.isJSONValueHitPolicy(payloads, subValue)
			if result.IsMatched {
				return result
			}
		}
	}
	return matchResult
}
