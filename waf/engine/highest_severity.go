package engine

import (
	"github.com/asalih/guardian/matches"
)

var HIGHEST_SEVERITY = "HIGHEST_SEVERITY"

func (t *TransactionMap) loadHighestSeverity() *TransactionMap {
	t.variableMap[HIGHEST_SEVERITY] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
