package engine

import (
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

func init() {

	TransactionMaps.variableMap["REQUEST_COOKIES_NAMES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			httpData := helpers.GetCookiesNames(executer.transaction.Request.Cookies())

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
}
