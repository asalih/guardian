package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadRequestLine() *TransactionMap {
	t.variableMap["REQUEST_LINE"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			uri := executer.transaction.Request.Host

			if executer.transaction.Request.TLS == nil {
				uri = "http://" + uri
			} else {
				uri = "https://" + uri
			}

			if executer.transaction.Request.RequestURI != "" {
				uri += executer.transaction.Request.RequestURI
			}

			line := executer.transaction.Request.Method + " " + uri
			return executer.rule.ExecuteRule(line)
		}}

	return t
}
