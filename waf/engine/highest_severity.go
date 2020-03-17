package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadHighestSeverity() *TransactionMap {
	t.variableMap["HIGHEST_SEVERITY"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
