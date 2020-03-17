package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadDuration() *TransactionMap {
	t.variableMap["DURATION"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult().SetMatch(true)
		}}

	return t
}
