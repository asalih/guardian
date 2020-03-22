package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadResponseStatus() *TransactionMap {
	t.variableMap["RESPONSE_STATUS"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			return executer.rule.ExecuteRule(executer.transaction.Response.StatusCode)
		}}

	return t
}
