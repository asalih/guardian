package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadRequestBodyLength() *TransactionMap {
	t.variableMap["REQUEST_BODY_LENGTH"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			bodyBufferLen := len(executer.transaction.BodyProcessor.GetBodyBuffer())

			return executer.rule.ExecuteRule(bodyBufferLen)
		}}

	return t
}
