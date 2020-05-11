package transformations

import "strings"

func init() {
	TransformationMaps.funcMap["trimRight"] = func(variableData interface{}) interface{} {
		return strings.TrimRight(variableData.(string), " ")
	}
}
