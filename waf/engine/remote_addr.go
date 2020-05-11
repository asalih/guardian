package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["REMOTE_ADDR"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return executer.rule.ExecuteRule(executer.transaction.Request.RemoteAddr)
		}}
}
