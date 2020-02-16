package operators

import (
	"strings"
)

func (opMap *OperatorMap) loadBeginsWith() {
	opMap.funcMap["beginsWith"] = func(expression interface{}, variableData interface{}) bool {
		return strings.HasPrefix(variableData.(string), expression.(string))
	}
}
