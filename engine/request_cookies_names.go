package engine

import (
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

var REQUEST_COOKIES_NAMES = "REQUEST_COOKIES_NAMES"

func (t *TransactionMap) loadRequestCookiesNames() *TransactionMap {

	t.variableMap[REQUEST_COOKIES_NAMES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			httpData := helpers.GetCookiesNames(executer.request.Cookies())

			matchResult := matches.NewMatchResult()

			if executer.variable.LengthCheckForCollection {
				lenOfCookies := 0
				for _, key := range httpData {
					if executer.variable.ShouldPassCheck(key) {
						continue
					}

					lenOfCookies++
				}

				return executer.rule.ExecuteRule(lenOfCookies)
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
