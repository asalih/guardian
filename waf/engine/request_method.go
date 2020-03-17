package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadRequestMethod() *TransactionMap {
	t.variableMap["REQUEST_METHOD"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return executer.rule.ExecuteRule(executer.transaction.Request.Method)
		}}

	return t
}
