package operators

import (
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadRx() {
	opMap.funcMap["rx"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)
		isMatch, _ := helpers.IsMatch(expression.(string), variableData.(string))

		return matchResult.SetMatch(isMatch)
	}
}
