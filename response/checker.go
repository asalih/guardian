package response

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/asalih/guardian/waf/bodyprocessor"
	"github.com/asalih/guardian/waf/engine"

	"github.com/asalih/guardian/matches"

	"github.com/PaesslerAG/gval"

	"github.com/asalih/guardian/data"
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/models"
)

//Checker Response checker
type Checker struct {
	ResponseWriter http.ResponseWriter
	Transaction    *engine.Transaction
	Target         *models.Target

	result         *models.RuleExecutionResult
	firewallResult chan *matches.FirewallMatchResult
	startTime      time.Time
}

/*NewResponseChecker Request checker initializer*/
func NewResponseChecker(w http.ResponseWriter, t *engine.Transaction, resp *http.Response, target *models.Target) *Checker {
	t.Response = resp
	t.ResponseBodyProcessor = bodyprocessor.NewResponseBodyProcessor(resp)
	return &Checker{w, t, target, nil, nil, time.Now()}
}

/*Handle Request checker handler func*/
func (r *Checker) Handle() bool {

	if !r.Target.WAFEnabled {
		return false
	}

	result := r.handleWAFChecker(3)

	if result {
		return result
	}

	result = r.handleWAFChecker(4)

	if result {
		return result
	}

	return r.handleFirewallRuleChecker()
}

func (r *Checker) handleFirewallRuleChecker() bool {
	firewallChannel := make(chan bool, 1)
	db := &data.DBHelper{}

	go func() {
		var wg sync.WaitGroup

		firewallRules := db.GetResponseFirewallRules(r.Target.ID)
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
			"response": map[string]interface{}{
				"status":        r.Transaction.Response.Status,
				"statusCode":    r.Transaction.Response.StatusCode,
				"cookie":        helpers.CookiesToString(r.Transaction.Response.Cookies()),
				"header":        helpers.HeadersToString(r.Transaction.Response.Header),
				"contentLength": r.Transaction.Response.ContentLength,
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

	/*for i := range r.firewallResult {
		//Action: 0 is block
		//Action: 1 is allow
		if i.IsMatched && i.FirewallRule.Action == 0 ||
			!i.IsMatched && i.FirewallRule.Action == 1 {
			r.ResponseWriter.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", r.Transaction.Request.URL.Path)

			db.LogFirewallMatchResult(i, r.Target, r.Transaction.Request.RequestURI, true)

			return true
		}
	}*/

	return false
}

func (r *Checker) handleFirewallPayload(rule *models.FirewallRule, mapForFwRules map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	evalResult, everr := gval.Evaluate(rule.Expression, mapForFwRules)

	if everr != nil {
		fmt.Println(everr)
	}

	//r.firewallResult <- matches.NewFirewallMatchResult(rule, evalResult.(bool)).Time(r.time)
	r.firewallResult <- matches.NewFirewallMatchResult(evalResult.(bool))
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
