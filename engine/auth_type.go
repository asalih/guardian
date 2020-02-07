package engine

import (
	"github.com/asalih/guardian/matches"
)

var AUTH_TYPE = "AUTH_TYPE"

func (t *Transaction) loadAuthType() *Transaction {
	t.variableMap[AUTH_TYPE] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			headerValue := executer.request.Header.Get("Authorization")

			return executer.rule.ExecuteRule(headerValue)
		}}

	return t
}
