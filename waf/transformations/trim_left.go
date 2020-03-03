package transformations

import "strings"

func (transform *TransformationMap) loadTrimLeft() {
	transform.funcMap["trimLeft"] = func(variableData interface{}) interface{} {
		return strings.TrimLeft(variableData.(string), " ")
	}
}
