package operators

import (
	"strings"
)

func (opMap *OperatorMap) loadContains() {
	opMap.funcMap["contains"] = func(expression interface{}, variableData interface{}) bool {
		return strings.Contains(variableData.(string), expression.(string))
	}
}
