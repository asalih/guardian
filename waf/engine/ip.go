package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["IP"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			return matches.NewMatchResult()

		}}
}
