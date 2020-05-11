package engine

import (
	"github.com/asalih/guardian/matches"
)

func init() {
	TransactionMaps.variableMap["TX"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			matchResult := matches.NewMatchResult()
			httpData := executer.transaction.tx

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
				matchResult = executer.rule.ExecuteRule(value)

				if matchResult.IsMatched {
					return matchResult
				}
			}

			return matchResult
		}}
}
