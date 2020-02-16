package engine

import (
	"mime"

	"github.com/asalih/guardian/matches"
)

var REQUEST_BODY_TYPE = "REQUEST_BODY_TYPE"
var MULTIPART_BOUNDARY = "MULTIPART_BOUNDARY"
var MULTIPART_ERROR = "MULTIPART_ERROR"

func (t *TransactionMap) loadRequestBodyType() *TransactionMap {
	t.variableMap[REQUEST_BODY_TYPE] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			contentType := executer.request.Header.Get("Content-Type")
			mediaType, _, _ := mime.ParseMediaType(contentType)

			return executer.rule.ExecuteRule(mediaType)
		}}

	t.variableMap[MULTIPART_BOUNDARY] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			contentType := executer.request.Header.Get("Content-Type")
			_, mediaParams, _ := mime.ParseMediaType(contentType)

			return executer.rule.ExecuteRule(mediaParams["boundary"])
		}}

	t.variableMap[MULTIPART_ERROR] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			contentType := executer.request.Header.Get("Content-Type")
			_, _, err := mime.ParseMediaType(contentType)

			multiPartError := 0
			if err != nil && contentType != "" {
				multiPartError = 1
			}

			return executer.rule.ExecuteRule(multiPartError)
		}}

	return t
}
