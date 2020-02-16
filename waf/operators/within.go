package operators

import (
	"strings"
)

func (opMap *OperatorMap) loadWithin() {
	opMap.funcMap["loadWithin"] = func(expression interface{}, variableData interface{}) bool {
		return strings.Contains(expression.(string), variableData.(string))
	}
}
