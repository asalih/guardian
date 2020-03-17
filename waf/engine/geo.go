package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadGeo() *TransactionMap {
	t.variableMap["GEO"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
