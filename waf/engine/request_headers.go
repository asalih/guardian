package engine

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadRequestHeaders() *TransactionMap {
	t.variableMap["REQUEST_HEADERS"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()
			httpData := executer.transaction.Request.Header

			if executer.variable.LengthCheckForCollection {
				lenOfHeaders := 0
				for key := range httpData {
					if executer.variable.ShouldPassCheck(key) {
						continue
					}

					lenOfHeaders++
				}

				return executer.rule.ExecuteRule(lenOfHeaders)
			}

			for key, value := range httpData {
				if executer.variable.ShouldPassCheck(key) {
					continue
				}
				matchResult = executer.rule.ExecuteRule(strings.Join(value, ","))

				if matchResult.IsMatched {
					return matchResult
				}
			}

			return matchResult
		}}

	return t
}
