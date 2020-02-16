package engine

import (
	"github.com/asalih/guardian/matches"
)

var FILES_SIZES = "FILES_SIZES"

func (t *TransactionMap) loadFilesSizes() *TransactionMap {
	t.variableMap[FILES_SIZES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()

			muliErr := executer.request.ParseMultipartForm(1024 * 1024 * 4)

			if muliErr != nil {
				return matchResult.SetMatch(true)
			}

			files := executer.request.MultipartForm.File
			for _, headers := range files {
				for _, head := range headers {
					matchResult = executer.rule.ExecuteRule(head.Size)

					if matchResult.IsMatched {
						return matchResult.SetMatch(true)
					}
				}
			}

			return matchResult

		}}

	return t
}
