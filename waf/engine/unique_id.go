package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadUniqueID() *TransactionMap {
	t.variableMap["UNIQUE_ID"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//Not implemented yet - might not needed.
			return matches.NewMatchResult().SetMatch(true)
		}}

	return t
}
