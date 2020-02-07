package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadGe() {
	opMap.funcMap["ge"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false).SetMatch(variableData.(int) >= expression.(int))
	}
}
