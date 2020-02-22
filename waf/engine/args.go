package engine

import (
	"github.com/asalih/guardian/matches"
)

var ARGS = "ARGS"
var ARGS_GET = "ARGS_GET"
var ARGS_POST = "ARGS_POST"

func (t *TransactionMap) loadArgs() *TransactionMap {
	t.variableMap[ARGS] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, true, true)
			}
			return argsHandler(executer, true, true)
		}}

	t.variableMap[ARGS_GET] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, true, false)
			}
			return argsHandler(executer, true, false)
		}}

	t.variableMap[ARGS_POST] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, false, true)
			}
			return argsHandler(executer, false, true)
		}}

	return t
}

func argsHandler(executer *TransactionExecuterModel, executeGet bool, executePost bool) *matches.MatchResult {
	matchResult := matches.NewMatchResult()
	if executeGet {
		queries := executer.transaction.Request.URL.Query()
		for q := range queries {
			if executer.variable.ShouldPassCheck(q) {
				continue
			}

			for _, value := range queries[q] {
				matchResult = executer.rule.ExecuteRule(value)

				if matchResult.IsMatched {
					return matchResult
				}
			}
		}
	}

	if executePost {
		err := executer.transaction.Request.ParseForm()

		if err != nil {
			matchResult.SetMatch(true)
			return matchResult
		}

		form := executer.transaction.Request.Form

		for f := range form {
			if executer.variable.ShouldPassCheck(f) {
				continue
			}
			for _, value := range form[f] {
				matchResult = executer.rule.ExecuteRule(value)

				if matchResult.IsMatched {
					return matchResult
				}
			}
		}
	}

	return matchResult
}

func argsLengthHandler(executer *TransactionExecuterModel, executeGet bool, executePost bool) *matches.MatchResult {
	matchResult := matches.NewMatchResult()
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
		err := executer.transaction.Request.ParseForm()

		if err != nil {
			matchResult.SetMatch(true)
			return matchResult
		}

		form := executer.transaction.Request.Form

		for f := range form {
			if executer.variable.ShouldPassCheck(f) {
				continue
			}
			lengthOfParams++
		}
	}

	return executer.rule.ExecuteRule(lengthOfParams)
}
