package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadRequestBody() *TransactionMap {
	t.variableMap["REQUEST_BODY"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			bodyBuffer := executer.transaction.BodyProcessor.GetBodyBuffer()

			if bodyBuffer == nil {
				return matches.NewMatchResult()
			}

			return executer.rule.ExecuteRule(string(bodyBuffer))
		}}

	return t
}
