package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadMatchedVars() *TransactionMap {
	t.variableMap["MATCHED_VARS"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	t.variableMap["MATCHED_VARS_NAMES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
