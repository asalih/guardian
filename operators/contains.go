package operators

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadEq() {
	opMap.funcMap["contains"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)

		return matchResult.SetMatch(strings.Contains(variableData.(string), expression.(string)))
	}
}
