package engine

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

var REQUEST_FILENAME = "REQUEST_FILENAME"

func (t *TransactionMap) loadRequestFilename() *TransactionMap {
	t.variableMap[REQUEST_FILENAME] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			httpData := executer.transaction.Request.URL.Path
			if !strings.HasSuffix(httpData, "/") {
				httpData += "/"
			}

			return executer.rule.ExecuteRule(httpData)
		}}

	return t
}
