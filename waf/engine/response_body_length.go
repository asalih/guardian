package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["RESPONSE_BODY_LENGTH"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			bodyBufferLen := len(executer.transaction.ResponseBodyProcessor.GetBodyBuffer())

			return executer.rule.ExecuteRule(bodyBufferLen)
		}}
}
