package engine

import (
	"github.com/asalih/guardian/matches"
)

var FILES_NAMES = "FILES_NAMES"
var MULTIPART_FILENAME = "MULTIPART_FILENAME"

func (t *TransactionMap) loadFilesNames() *TransactionMap {
	transData := &TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
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
					return matchResult.SetMatch(true)
				}
			}
		}

		return matchResult

	}}

	t.variableMap[FILES_NAMES] = transData
	t.variableMap[MULTIPART_FILENAME] = transData

	return t
}
