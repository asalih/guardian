package engine

import (
	"github.com/asalih/guardian/matches"
)

var ARGS_COMBINED_SIZE = "ARGS_COMBINED_SIZE"

func (t *TransactionMap) loadArgsCombinedSize() *TransactionMap {
	t.variableMap[ARGS_COMBINED_SIZE] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()

			sizeOfParams := 0

			queries := executer.transaction.Request.URL.Query()
			for q := range queries {
				if executer.variable.ShouldPassCheck(q) {
					continue
				}

				sizeOfParams += len(queries[q])
			}

			err := executer.transaction.Request.ParseForm()

			if err != nil {
				matchResult.SetMatch(true)
				return matchResult
			}

			form := executer.transaction.Request.Form

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
