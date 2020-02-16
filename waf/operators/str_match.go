package operators

import (
	"strings"
)

func (opMap *OperatorMap) loadStrmatch() {
	opMap.funcMap["strmatch"] = func(expression interface{}, variableData interface{}) bool {
		return strings.Contains(variableData.(string), expression.(string))
	}
}
