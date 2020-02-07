package operators

import (
	"github.com/asalih/guardian/matches"
	"github.com/koangel/grapeSQLI"
)

func (opMap *OperatorMap) loadDetectXss() {
	opMap.funcMap["detectXSS"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)
		matchResult.SetMatch(GSQLI.XSSParser(variableData.(string)))

		return matchResult
	}
}
