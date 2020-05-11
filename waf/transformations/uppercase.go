package transformations

import "strings"

func init() {
	TransformationMaps.funcMap["uppercase"] = func(variableData interface{}) interface{} {
		return strings.ToUpper(variableData.(string))
	}
}
