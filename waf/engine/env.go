package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadEnv() *TransactionMap {
	t.variableMap["ENV"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult().SetMatch(false)
		}}

	return t
}
