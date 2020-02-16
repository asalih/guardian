package engine

import (
	"github.com/asalih/guardian/matches"
)

var INBOUND_DATA_ERROR = "INBOUND_DATA_ERROR"

func (t *TransactionMap) loadInboundDataError() *TransactionMap {
	t.variableMap[INBOUND_DATA_ERROR] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
