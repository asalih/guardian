package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadContains() {
	opMap.funcMap["eq"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false).SetMatch(expression.(int) == variableData.(int))
	}
}
