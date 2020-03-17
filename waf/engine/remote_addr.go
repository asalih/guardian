package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadRemoteAddr() *TransactionMap {
	t.variableMap["REMOTE_ADDR"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return executer.rule.ExecuteRule(executer.transaction.Request.RemoteAddr)
		}}

	return t
}
