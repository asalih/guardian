package engine

import (
	"github.com/asalih/guardian/matches"
)

var REQUEST_COOKIES = "REQUEST_COOKIES"

func (t *Transaction) loadRequestCookies() *Transaction {

	t.variableMap[REQUEST_COOKIES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult(false)
			httpData := executer.request.Cookies()

			if executer.variable.LengthCheckForCollection {

				lenOfCookies := 0
				for _, cookie := range httpData {
					if executer.variable.ShouldPassCheck(cookie.Name) {
						continue
					}

					lenOfCookies++
				}

				return executer.rule.ExecuteRule(lenOfCookies)
			}

			for _, cookie := range httpData {
				if executer.variable.ShouldPassCheck(cookie.Name) {
					continue
				}

				matchResult = executer.rule.ExecuteRule(cookie.Value)

				if matchResult.IsMatched {
					return matchResult
				}
			}

			return matchResult
		}}

	return t
}
