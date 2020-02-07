package engine

import (
	"github.com/asalih/guardian/matches"
)

var FILES = "FILES"

func (t *Transaction) loadFiles() *Transaction {
	t.variableMap[FILES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult(false)

			muliErr := executer.request.ParseMultipartForm(1024 * 1024 * 4)

			if muliErr != nil {
				return matchResult.SetMatch(true)
			}

			files := executer.request.MultipartForm.File

			for _, headers := range files {
				for _, head := range headers {

					matchResult = executer.rule.ExecuteRule(head.Filename)

					if matchResult.IsMatched {
						return matchResult
					}
				}
			}

			return matchResult
		}}

	return t
}
