package operators

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadWithin() {
	opMap.funcMap["loadWithin"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)

		return matchResult.SetMatch(strings.Contains(expression.(string), variableData.(string)))
	}
}
