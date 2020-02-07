package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadNoMatch() {
	opMap.funcMap["noMatch"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false)
	}
}
