package engine

import (
	"github.com/asalih/guardian/matches"
	"github.com/asalih/guardian/waf/bodyprocessor"
)

func (t *TransactionMap) loadFilesCombinedSize() *TransactionMap {
	t.variableMap["FILES_COMBINED_SIZE"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()

			switch executer.transaction.RequestBodyProcessor.(type) {
			case *bodyprocessor.MultipartProcessor:
				files := executer.transaction.Request.MultipartForm.File
				totalSize := int64(0)
				for _, headers := range files {
					for _, head := range headers {
						totalSize += head.Size
					}
				}

				return executer.rule.ExecuteRule(totalSize)
			}

			return matchResult.SetMatch(false)

		}}

	return t
}
