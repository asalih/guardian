package engine

import (
	"github.com/asalih/guardian/matches"
)

var REQUEST_URI = "REQUEST_URI"
var REQUEST_PROTOCOL = "REQUEST_PROTOCOL"
var REQUEST_URI_RAW = "REQUEST_URI_RAW"
var REQUEST_BASENAME = "REQUEST_BASENAME"

func (t *Transaction) loadRequestUri() *Transaction {

	t.variableMap[REQUEST_URI] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			return executer.rule.ExecuteRule(executer.request.RequestURI)
		}}

	t.variableMap[REQUEST_PROTOCOL] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			proto := "http"
			if executer.request.TLS != nil {
				proto = "https"
			}

			return executer.rule.ExecuteRule(proto)
		}}

	t.variableMap[REQUEST_URI_RAW] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			uriRaw := ""
			if executer.request.TLS != nil {
				uriRaw = "https://" + executer.request.Host + executer.request.RequestURI
			} else {
				uriRaw = "http://" + executer.request.Host + executer.request.RequestURI
			}

			return executer.rule.ExecuteRule(uriRaw)
		}}

	t.variableMap[REQUEST_BASENAME] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			return executer.rule.ExecuteRule(executer.request.URL.Path)
		}}

	return t
}
