package engine

import (
	"mime"

	"github.com/asalih/guardian/matches"
)

func inşt() {
	TransactionMaps.variableMap["REQUEST_BODY_TYPE"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			contentType := executer.transaction.Request.Header.Get("Content-Type")
			mediaType, _, _ := mime.ParseMediaType(contentType)

			return executer.rule.ExecuteRule(mediaType)
		}}

	TransactionMaps.variableMap["MULTIPART_BOUNDARY"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			contentType := executer.transaction.Request.Header.Get("Content-Type")
			_, mediaParams, _ := mime.ParseMediaType(contentType)

			return executer.rule.ExecuteRule(mediaParams["boundary"])
		}}

	TransactionMaps.variableMap["MULTIPART_ERROR"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			contentType := executer.transaction.Request.Header.Get("Content-Type")
			_, _, err := mime.ParseMediaType(contentType)

			multiPartError := 0
			if err != nil && contentType != "" {
				multiPartError = 1
			}

			return executer.rule.ExecuteRule(multiPartError)
		}}
}
