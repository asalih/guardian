package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadUnconditionalMatch() {
	opMap.funcMap["unconditionalMatch"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(true)
	}
}
