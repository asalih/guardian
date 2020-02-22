package request

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/asalih/guardian/matches"

	"github.com/PaesslerAG/gval"

	"github.com/asalih/guardian/data"
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/waf/engine"
)

var staticSuffix = []string{".js", ".css", ".png", ".jpg", ".gif", ".bmp", ".svg", ".ico"}

/*Checker Cheks the requests init*/
type Checker struct {
	ResponseWriter http.ResponseWriter
	Target         *models.Target
	Transaction    *engine.Transaction

	result         *models.RuleExecutionResult
	firewallResult chan *matches.FirewallMatchResult
	startTime      time.Time
}

/*NewRequestChecker Request checker initializer*/
func NewRequestChecker(w http.ResponseWriter, r *http.Request, target *models.Target) *Checker {
	return &Checker{w, target, engine.NewTransaction(r), nil, nil, time.Now()}
}

/*Handle Request checker handler func*/
func (r *Checker) Handle() bool {

	if !r.Target.WAFEnabled || r.Transaction.Request.Method == "GET" && r.IsStaticResource(r.Transaction.Request.URL.Path) {
		return false
	}

	result := r.handleWAFChecker(1)

	if result {
		return result
	}

	result = r.handleWAFChecker(2)

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
				"src": r.Transaction.Request.RemoteAddr,
			},
			"http": map[string]interface{}{
				"query":    r.Transaction.Request.URL.RawQuery,
				"path":     r.Transaction.Request.URL.Path,
				"host":     r.Transaction.Request.URL.Host,
				"cookie":   helpers.CookiesToString(r.Transaction.Request.Cookies()),
				"header":   helpers.HeadersToString(r.Transaction.Request.Header),
				"method":   r.Transaction.Request.Method,
				"protocol": r.Transaction.Request.Proto,
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

func (r *Checker) handleWAFChecker(phase int) bool {

	done := make(chan bool, 1)

	go func() {

		for _, rule := range models.RulesCollection[phase] {

			//ruleStartTime := time.Now()
			matchResult := r.Transaction.Execute(rule)

			if matchResult == nil {
				continue
			}

			if matchResult.IsMatched && rule.ShouldBlock() {
				r.result = &models.RuleExecutionResult{matchResult, rule}
				break
			} else if !matchResult.IsMatched && !matchResult.DefaultState && !rule.ShouldBlock() {
				matchResult.SetMatch(true)
				r.result = &models.RuleExecutionResult{matchResult, rule}
				break
			}

			//Line 127 and below commented lines are for calculating each rulr exec time
			//passed := helpers.CalcTime(ruleStartTime, time.Now())
			//if passed > 0 {
			//	fmt.Println(rule.Action.ID + " took " + strconv.FormatInt(passed, 10) + " ms.")
			//}
		}

		done <- true
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Minute):
		panic("failed to execute rules.")
	}

	if r.result != nil && r.result.MatchResult.IsMatched {
		r.ResponseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", url.QueryEscape(r.Transaction.Request.URL.Path))

		if r.result.Rule.Action.LogAction == models.LogActionLog {
			db := &data.DBHelper{}
			go db.LogMatchResult(r.result, "TEMP", r.Target, r.Transaction.Request.RequestURI, false)
		}

		return true
	}

	return false
}

func (r *Checker) handleFirewallPayload(rule *models.FirewallRule, mapForFwRules map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	evalResult, everr := gval.Evaluate(rule.Expression, mapForFwRules)

	if everr != nil {
		fmt.Println(everr)
	}

	r.firewallResult <- matches.NewFirewallMatchResult(evalResult.(bool))
}
