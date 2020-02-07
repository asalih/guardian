package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadStreq() {
	opMap.funcMap["streq"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false).SetMatch(expression.(string) == variableData.(string))
	}
}
