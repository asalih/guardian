package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadXML() *TransactionMap {
	t.variableMap["XML"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return matches.NewMatchResult()

		}}

	return t
}
