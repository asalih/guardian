package operators

import (
	"strings"
)

func inir() {
	OperatorMaps.funcMap["within"] = func(expression interface{}, variableData interface{}) bool {
		return strings.Contains(expression.(string), variableData.(string))
	}
}
