package engine

import (
	"github.com/asalih/guardian/matches"
)

var MATCHED_VARS = "MATCHED_VARS"
var MATCHED_VARS_NAMES = "MATCHED_VARS_NAMES"

func (t *Transaction) loadMatchedVars() *Transaction {
	t.variableMap[MATCHED_VARS] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	t.variableMap[MATCHED_VARS_NAMES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	return t
}
