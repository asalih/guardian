package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadLt() {
	opMap.funcMap["lt"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false).SetMatch(variableData.(int) < expression.(int))
	}
}
