package engine

import (
	"github.com/asalih/guardian/matches"
)

var DURATION = "DURATION"

func (t *Transaction) loadDuration() *Transaction {
	t.variableMap[DURATION] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	return t
}
