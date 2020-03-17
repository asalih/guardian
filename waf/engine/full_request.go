package engine

import (
	"fmt"
	"strings"

	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadFullRequestAndLength() *TransactionMap {
	t.variableMap["FULL_REQUEST"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			httpData := formatRequest(executer.transaction)

			return executer.rule.ExecuteRule(httpData)
		}}

	t.variableMap["FULL_REQUEST_LENGTH"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			httpData := len(formatRequest(executer.transaction))

			return executer.rule.ExecuteRule(httpData)
		}}

	return t
}

// formatRequest generates ascii representation of a request
func formatRequest(t *Transaction) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", t.Request.Method, t.Request.URL, t.Request.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", t.Request.Host))
	// Loop through headers
	for name, headers := range t.Request.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if t.Request.Method == "POST" || t.Request.Method == "PUT" {
		t.BodyProcessor.GetBody()
		request = append(request, "\n")
		request = append(request, t.Request.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
