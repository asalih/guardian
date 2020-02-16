package engine

import (
	"github.com/asalih/guardian/matches"
)

var QUERY_STRING = "QUERY_STRING"

func (t *TransactionMap) loadQueryString() *TransactionMap {
	t.variableMap[QUERY_STRING] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return executer.rule.ExecuteRule(executer.request.URL.RawQuery)
		}}

	return t
}
