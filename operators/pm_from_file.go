package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadPmFromFile() {
	fn := func(expression interface{}, variableData interface{}) *matches.MatchResult {
		//TODO: might have to review
		return matches.NewMatchResult(false)
	}

	opMap.funcMap["pmf"] = fn
	opMap.funcMap["pmFromFile"] = fn
}
