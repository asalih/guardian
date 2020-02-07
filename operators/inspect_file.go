package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadInspectFile() {
	opMap.funcMap["inspectFile"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		//TODO: might have to review
		return matches.NewMatchResult(false)
	}
}
