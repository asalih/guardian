package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadIPMatchFromFile() {
	fn := func(expression interface{}, variableData interface{}) *matches.MatchResult {
		//TODO: might have to review
		return matches.NewMatchResult(false)
	}

	opMap.funcMap["ipMatchF"] = fn
	opMap.funcMap["ipMatchFromFile"] = fn
}
