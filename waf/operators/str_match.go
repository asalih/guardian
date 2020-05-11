package operators

import (
	"strings"
)

func init() {
	OperatorMaps.funcMap["strmatch"] = func(expression interface{}, variableData interface{}) bool {
		return strings.Contains(variableData.(string), expression.(string))
	}
}
