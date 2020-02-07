package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadGt() {
	opMap.funcMap["gt"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false).SetMatch(variableData.(int) > expression.(int))
	}
}
