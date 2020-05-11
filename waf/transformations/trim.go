package transformations

import "strings"

func init() {
	TransformationMaps.funcMap["trim"] = func(variableData interface{}) interface{} {
		return strings.Trim(variableData.(string), " ")
	}
}
