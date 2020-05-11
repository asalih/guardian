package operators

import (
	"strings"
)

func init() {
	OperatorMaps.funcMap["contains"] = func(expression interface{}, variableData interface{}) bool {
		return strings.Contains(variableData.(string), expression.(string))
	}
}
