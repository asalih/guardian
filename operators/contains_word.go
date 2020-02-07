package operators

import (
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadContainsWord() {
	opMap.funcMap["containsWord"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)
		isMatch, _ := helpers.IsMatch("\b"+expression.(string)+"\b", variableData.(string))

		return matchResult.SetMatch(isMatch)
	}
}
