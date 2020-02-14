package operators

import (
	"strings"
)

func (opMap *OperatorMap) loadEndsWith() {
	opMap.funcMap["endsWith"] = func(expression interface{}, variableData interface{}) bool {

		return strings.HasSuffix(variableData.(string), expression.(string))
	}
}
