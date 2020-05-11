package engine

import (
	"time"

	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["DURATION"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			duration := helpers.CalcTime(executer.transaction.duration, time.Now())

			return executer.rule.ExecuteRule(duration)
		}}
}
