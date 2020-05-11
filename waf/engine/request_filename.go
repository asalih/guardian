package engine

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["REQUEST_FILENAME"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			httpData := executer.transaction.Request.URL.Path
			if !strings.HasSuffix(httpData, "/") {
				httpData += "/"
			}

			return executer.rule.ExecuteRule(httpData)
		}}
}
