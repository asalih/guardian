package engine

import (
	"github.com/asalih/guardian/matches"
)

var FILES = "FILES"

func (t *TransactionMap) loadFiles() *TransactionMap {
	t.variableMap[FILES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()

			muliErr := executer.transaction.Request.ParseMultipartForm(1024 * 1024 * 4)

			if muliErr != nil {
				return matchResult.SetMatch(false)
			}

			files := executer.transaction.Request.MultipartForm.File

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
