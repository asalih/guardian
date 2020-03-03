package transformations

import "strings"

func (transform *TransformationMap) loadUppercase() {
	transform.funcMap["uppercase"] = func(variableData interface{}) interface{} {
		return strings.ToUpper(variableData.(string))
	}
}
