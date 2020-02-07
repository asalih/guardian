package engine

import (
	"github.com/asalih/guardian/matches"
)

var ENV = "ENV"

func (t *Transaction) loadEnv() *Transaction {
	t.variableMap[ENV] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	return t
}
