package engine

import (
	"github.com/asalih/guardian/matches"
)

var ARGS_NAMES = "ARGS_NAMES"
var ARGS_GET_NAMES = "ARGS_GET_NAMES"
var ARGS_POST_NAMES = "ARGS_POST_NAMES"

func (t *Transaction) loadArgsNames() *Transaction {
	t.variableMap[ARGS_NAMES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, true, true)
			}
			return argsHandler(executer, true, true)
		}}

	t.variableMap[ARGS_GET_NAMES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, true, false)
			}
			return argsHandler(executer, true, false)
		}}

	t.variableMap[ARGS_POST_NAMES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			if executer.variable.LengthCheckForCollection {
				return argsLengthHandler(executer, false, true)
			}
			return argsHandler(executer, false, true)
		}}

	return t
}

func argsNameHandler(executer *TransactionExecuterModel, executeGet bool, executePost bool) *matches.MatchResult {
	matchResult := matches.NewMatchResult(false)
	if executeGet {
		queries := executer.request.URL.Query()
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
		err := executer.request.ParseForm()

		if err != nil {
			matchResult.SetMatch(true)
			return matchResult
		}

		form := executer.request.Form

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
	matchResult := matches.NewMatchResult(false)
	lengthOfParams := 0
	if executeGet {
		queries := executer.request.URL.Query()
		for q := range queries {
			if executer.variable.ShouldPassCheck(q) {
				continue
			}
			lengthOfParams++
		}
	}

	if executePost {
		err := executer.request.ParseForm()

		if err != nil {
			matchResult.SetMatch(true)
			return matchResult
		}

		form := executer.request.Form

		for f := range form {
			if executer.variable.ShouldPassCheck(f) {
				continue
			}
			lengthOfParams++
		}
	}

	return executer.rule.ExecuteRule(lengthOfParams)
}
