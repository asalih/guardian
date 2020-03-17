package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadArgsNames() *TransactionMap {
	t.variableMap["ARGS_NAMES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, true, true)
			}
			return argsHandler(executer, true, true)
		}}

	t.variableMap["ARGS_GET_NAMES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, true, false)
			}
			return argsHandler(executer, true, false)
		}}

	t.variableMap["ARGS_POST_NAMES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, false, true)
			}
			return argsHandler(executer, false, true)
		}}

	return t
}

func argsNameHandler(executer *TransactionExecuterModel, executeGet bool, executePost bool) *matches.MatchResult {
	matchResult := matches.NewMatchResult()
	if executeGet {
		queries := executer.transaction.Request.URL.Query()
		for q := range queries {
			if executer.variable.ShouldPassCheck(q) {
				continue
			}
			matchResult = executer.rule.ExecuteRule(q)

			if matchResult.IsMatched {
				return matchResult
			}
		}
	}

	if executePost {

		form := executer.transaction.Request.Form

		for f := range form {
			if executer.variable.ShouldPassCheck(f) {
				continue
			}
			matchResult = executer.rule.ExecuteRule(f)

			if matchResult.IsMatched {
				return matchResult
			}
		}
	}

	return matchResult
}

func argsNameLengthHandler(executer *TransactionExecuterModel, executeGet bool, executePost bool) *matches.MatchResult {

	lengthOfParams := 0
	if executeGet {
		queries := executer.transaction.Request.URL.Query()
		for q := range queries {
			if executer.variable.ShouldPassCheck(q) {
				continue
			}
			lengthOfParams++
		}
	}

	if executePost {

		form := executer.transaction.BodyProcessor.GetBody()

		for f := range form {
			if executer.variable.ShouldPassCheck(f) {
				continue
			}
			lengthOfParams++
		}
	}

	return executer.rule.ExecuteRule(lengthOfParams)
}
