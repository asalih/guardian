package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["QUERY_STRING"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return executer.rule.ExecuteRule(executer.transaction.Request.URL.RawQuery)
		}}
}
