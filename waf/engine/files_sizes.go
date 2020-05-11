package engine

import (
	"github.com/asalih/guardian/matches"
	"github.com/asalih/guardian/waf/bodyprocessor"
)

func init() {
	TransactionMaps.variableMap["FILES_SIZES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()

			switch executer.transaction.RequestBodyProcessor.(type) {
			case *bodyprocessor.MultipartProcessor:

				files := executer.transaction.Request.MultipartForm.File
				for _, headers := range files {
					for _, head := range headers {
						matchResult = executer.rule.ExecuteRule(head.Size)

						if matchResult.IsMatched {
							return matchResult.SetMatch(true)
						}
					}
				}

				return matchResult
			}

			return matchResult.SetMatch(false)
		}}
}
