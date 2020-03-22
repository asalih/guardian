package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadResponseBodyLength() *TransactionMap {
	t.variableMap["RESPONSE_BODY_LENGTH"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			bodyBufferLen := len(executer.transaction.ResponseBodyProcessor.GetBodyBuffer())

			return executer.rule.ExecuteRule(bodyBufferLen)
		}}

	return t
}
