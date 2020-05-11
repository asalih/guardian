package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["AUTH_TYPE"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			headerValue := executer.transaction.Request.Header.Get("Authorization")

			return executer.rule.ExecuteRule(headerValue)
		}}
}
