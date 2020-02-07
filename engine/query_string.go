package engine

import (
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

var QUERY_STRING = "QUERY_STRING"

func (t *Transaction) loadQueryString() *Transaction {
	t.variableMap[QUERY_STRING] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			httpData := helpers.UnEscapeRawValue(executer.request.URL.RawQuery)

			return executer.rule.ExecuteRule(httpData)
		}}

	return t
}
