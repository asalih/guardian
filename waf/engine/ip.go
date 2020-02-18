package engine

import (
	"github.com/asalih/guardian/matches"
)

var IP = "IP"

func (t *TransactionMap) loadIP() *TransactionMap {
	t.variableMap[IP] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return matches.NewMatchResult()

		}}

	return t
}
