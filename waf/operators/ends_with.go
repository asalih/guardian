package operators

import (
	"strings"
)

func init() {
	OperatorMaps.funcMap["endsWith"] = func(expression interface{}, variableData interface{}) bool {

		return strings.HasSuffix(variableData.(string), expression.(string))
	}
}
