package engine

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/asalih/guardian/matches"
)

var FULL_REQUEST = "FULL_REQUEST"
var FULL_REQUEST_LENGTH = "FULL_REQUEST_LENGTH"

func (t *TransactionMap) loadFullRequestAndLength() *TransactionMap {
	t.variableMap[FULL_REQUEST] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			httpData := formatRequest(executer.request)

			return executer.rule.ExecuteRule(httpData)
		}}

	t.variableMap[FULL_REQUEST_LENGTH] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			httpData := len(formatRequest(executer.request))

			return executer.rule.ExecuteRule(httpData)
		}}

	return t
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" || r.Method == "PUT" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
