package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadQueryString() *TransactionMap {
	t.variableMap["QUERY_STRING"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return executer.rule.ExecuteRule(executer.transaction.Request.URL.RawQuery)
		}}

	return t
}
