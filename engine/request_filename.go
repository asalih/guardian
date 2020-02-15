package engine

import (
	"strings"

	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

var REQUEST_FILENAME = "REQUEST_FILENAME"

func (t *TransactionMap) loadRequestFilename() *TransactionMap {
	t.variableMap[REQUEST_FILENAME] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			httpData := helpers.UnEscapeRawValue(executer.request.URL.Path)
			if !strings.HasSuffix(httpData, "/") {
				httpData += "/"
			}

			return executer.rule.ExecuteRule(httpData)
		}}

	return t
}
