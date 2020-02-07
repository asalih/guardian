package engine

import (
	"github.com/asalih/guardian/matches"
)

var GEO = "GEO"

func (t *Transaction) loadGeo() *Transaction {
	t.variableMap[GEO] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	return t
}
