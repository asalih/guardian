package engine

import (
	"github.com/asalih/guardian/matches"
)

var FILES_NAMES = "FILES_NAMES"
var MULTIPART_FILENAME = "MULTIPART_FILENAME"

func (t *Transaction) loadFilesNames() *Transaction {
	transData := &TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
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
