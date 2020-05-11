package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {

	TransactionMaps.variableMap["REQUEST_COOKIES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()
			httpData := executer.transaction.Request.Cookies()

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
}
