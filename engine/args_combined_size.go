package engine

import (
	"github.com/asalih/guardian/matches"
)

var ARGS_COMBINED_SIZE = "ARGS_COMBINED_SIZE"

func (t *Transaction) loadArgsCombinedSize() *Transaction {
	t.variableMap[ARGS_COMBINED_SIZE] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult(false)

			sizeOfParams := 0

			queries := executer.request.URL.Query()
			for q := range queries {
				if executer.variable.ShouldPassCheck(q) {
					continue
				}

				sizeOfParams += len(queries[q])
			}

			err := executer.request.ParseForm()

			if err != nil {
				matchResult.SetMatch(true)
				return matchResult
			}

			form := executer.request.Form

			for f := range form {
				if executer.variable.ShouldPassCheck(f) {
					continue
				}

				sizeOfParams += len(form[f])
			}

			return executer.rule.ExecuteRule(sizeOfParams)
		}}

	return t
}
