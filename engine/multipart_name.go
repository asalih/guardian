package engine

import (
	"github.com/asalih/guardian/matches"
)

var MULTIPART_NAME = "MULTIPART_NAME"

func (t *Transaction) loadMultipartName() *Transaction {
	transData := &TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)

		muliErr := executer.request.ParseMultipartForm(1024 * 1024 * 4)

		if muliErr != nil {
			return matchResult.SetMatch(true)
		}

		files := executer.request.MultipartForm.File
		for key, _ := range files {

			matchResult = executer.rule.ExecuteRule(key)

			if matchResult.IsMatched {
				return matchResult.SetMatch(true)
			}

		}

		return matchResult

	}}

	t.variableMap[FILES_NAMES] = transData
	t.variableMap[MULTIPART_FILENAME] = transData

	return t
}
