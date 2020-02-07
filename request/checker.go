package request

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/asalih/guardian/matches"

	"github.com/PaesslerAG/gval"

	"github.com/asalih/guardian/data"
	"github.com/asalih/guardian/engine"
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/models"
)

var staticSuffix = []string{".js", ".css", ".png", ".jpg", ".gif", ".bmp", ".svg", ".ico"}

/*Checker Cheks the requests init*/
type Checker struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Target         *models.Target
	Transaction    *engine.Transaction

	result         *matches.MatchResult
	firewallResult chan *matches.FirewallMatchResult
	startTime      time.Time
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

	r.Transaction = engine.NewTransaction(r.Request)
	result := r.handleWAFChecker()

	if result {
		return result
	}

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

		r.firewallResult = make(chan *matches.FirewallMatchResult, lenOfRules)

		wg.Add(lenOfRules)

		mapForFwRules := map[string]interface{}{
			"ip": map[string]interface{}{
				"src": r.Request.RemoteAddr,
			},
			"http": map[string]interface{}{
				"query":    r.Request.URL.RawQuery,
				"path":     r.Request.URL.Path,
				"host":     r.Request.URL.Host,
				"cookie":   helpers.CookiesToString(r.Request.Cookies()),
				"header":   helpers.HeadersToString(r.Request.Header),
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

	return false
}

func (r *Checker) handleWAFChecker() bool {

	done := make(chan bool, 1)

	go func() {

		for _, rule := range models.RulesCollection {
			matchedRules := r.Transaction.ExecuteRule(rule)

			if matchedRules.IsMatched {
				r.result = matchedRules
				break
			}
		}

		done <- true
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Minute):
		panic("failed to execute rules.")
	}

	if r.result != nil && r.result.IsMatched {
		//if m.Action == "block" {
		r.ResponseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", r.Request.URL.Path)

		db := &data.DBHelper{}

		go db.LogMatchResult(r.result, "TEMP", r.Target, r.Request.RequestURI, false)

		return true
		//}
	}

	return false
}

func (r *Checker) handleFirewallPayload(rule *models.FirewallRule, mapForFwRules map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	evalResult, everr := gval.Evaluate(rule.Expression, mapForFwRules)

	if everr != nil {
		fmt.Println(everr)
	}

	//r.firewallResult <- models.NewFirewallMatchResult(rule, evalResult.(bool)).Time(r.time)
	r.firewallResult <- matches.NewFirewallMatchResult(evalResult.(bool))
}
