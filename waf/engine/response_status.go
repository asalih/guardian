package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["RESPONSE_STATUS"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			return executer.rule.ExecuteRule(executer.transaction.Response.StatusCode)
		}}
}
