package transformations

import "strings"

func (transform *TransformationMap) loadTrim() {
	transform.funcMap["trim"] = func(variableData interface{}) interface{} {
		return strings.Trim(variableData.(string), " ")
	}
}
