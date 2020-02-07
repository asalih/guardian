package engine

import (
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

var REQUEST_HEADERS_NAMES = "REQUEST_HEADERS_NAMES"

func (t *Transaction) loadRequestHeadersNames() *Transaction {
	t.variableMap[REQUEST_HEADERS_NAMES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			httpData := helpers.GetHeadersNames(executer.request.Header)

			matchResult := matches.NewMatchResult(false)

			if executer.variable.LengthCheckForCollection {
				lenOfHeaders := 0
				for _, key := range httpData {
					if executer.variable.ShouldPassCheck(key) {
						continue
					}

					lenOfHeaders++
				}

				return executer.rule.ExecuteRule(lenOfHeaders)
			}

			for _, key := range httpData {
				if executer.variable.ShouldPassCheck(key) {
					continue
				}
				matchResult = executer.rule.ExecuteRule(key)

				if matchResult.IsMatched {
					return matchResult
				}
			}

			return matchResult
		}}

	return t
}
