package operators

import (
	"strings"
)

func init() {
	OperatorMaps.funcMap["beginsWith"] = func(expression interface{}, variableData interface{}) bool {
		return strings.HasPrefix(variableData.(string), expression.(string))
	}
}
