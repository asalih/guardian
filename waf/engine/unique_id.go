package engine

import (
	"github.com/asalih/guardian/matches"
)

var UNIQUE_ID = "UNIQUE_ID"

func (t *TransactionMap) loadUniqueID() *TransactionMap {
	t.variableMap[UNIQUE_ID] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return matches.NewMatchResult().SetMatch(true)
		}}

	return t
}
