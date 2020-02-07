package engine

import (
	"github.com/asalih/guardian/matches"
)

var MATCHED_VAR = "MATCHED_VAR"
var MATCHED_VAR_NAME = "MATCHED_VAR_NAME"

func (t *Transaction) loadMatchedVar() *Transaction {
	t.variableMap[MATCHED_VAR] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	t.variableMap[MATCHED_VAR_NAME] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	return t
}
