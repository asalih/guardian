package transformations

import "strings"

func (transform *TransformationMap) loadTrimRight() {
	transform.funcMap["trimRight"] = func(variableData interface{}) interface{} {
		return strings.TrimRight(variableData.(string), " ")
	}
}
