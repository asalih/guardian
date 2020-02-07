package operators

import (
	"github.com/asalih/guardian/matches"
	"github.com/koangel/grapeSQLI"
)

func (opMap *OperatorMap) loadDetectSqli() {
	opMap.funcMap["detectSQLi"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)

		if err := GSQLI.SQLInject(variableData.(string)); err != nil {
			matchResult.SetMatch(true)
		}

		return matchResult
	}
}
