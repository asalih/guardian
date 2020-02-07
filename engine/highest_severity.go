package engine

import (
	"github.com/asalih/guardian/matches"
)

var HIGHEST_SEVERITY = "HIGHEST_SEVERITY"

func (t *Transaction) loadHighestSeverity() *Transaction {
	t.variableMap[HIGHEST_SEVERITY] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	return t
}
