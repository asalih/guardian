package engine

import (
	"time"

	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadDuration() *TransactionMap {
	t.variableMap["DURATION"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			duration := helpers.CalcTime(executer.transaction.duration, time.Now())

			return executer.rule.ExecuteRule(duration)
		}}

	return t
}
