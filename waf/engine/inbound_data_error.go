package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadInboundDataError() *TransactionMap {
	t.variableMap["INBOUND_DATA_ERROR"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
