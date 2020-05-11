package transformations

import "strings"

func init() {
	TransformationMaps.funcMap["trimLeft"] = func(variableData interface{}) interface{} {
		return strings.TrimLeft(variableData.(string), " ")
	}
}
