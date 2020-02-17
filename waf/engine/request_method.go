package engine

import (
	"github.com/asalih/guardian/matches"
)

var REQUEST_METHOD = "REQUEST_METHOD"

func (t *TransactionMap) loadRequestMethod() *TransactionMap {
	t.variableMap[REQUEST_METHOD] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return executer.rule.ExecuteRule(executer.transaction.request.Method)
		}}

	return t
}
