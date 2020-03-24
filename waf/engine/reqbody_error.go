package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadReqBodyError() *TransactionMap {
	t.variableMap["REQBODY_ERROR"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.transaction.RequestBodyProcessor.HasBodyError() {
				return executer.rule.ExecuteRule("1")
			}

			return executer.rule.ExecuteRule("0")
		}}

	return t
}
