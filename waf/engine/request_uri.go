package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadRequestURI() *TransactionMap {

	t.variableMap["REQUEST_URI"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			return executer.rule.ExecuteRule(executer.transaction.Request.RequestURI)
		}}

	t.variableMap["REQUEST_PROTOCOL"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			proto := "http"
			if executer.transaction.Request.TLS != nil {
				proto = "https"
			}

			return executer.rule.ExecuteRule(proto)
		}}

	t.variableMap["REQUEST_URI_RAW"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			uriRaw := ""
			if executer.transaction.Request.TLS != nil {
				uriRaw = "https://" + executer.transaction.Request.Host + executer.transaction.Request.RequestURI
			} else {
				uriRaw = "http://" + executer.transaction.Request.Host + executer.transaction.Request.RequestURI
			}

			return executer.rule.ExecuteRule(uriRaw)
		}}

	t.variableMap["REQUEST_BASENAME"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			return executer.rule.ExecuteRule(executer.transaction.Request.URL.Path)
		}}

	return t
}
