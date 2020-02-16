package transformations

import (
	"strings"
)

func (transform *TransformationMap) loadLowercase() {
	transform.funcMap["lowercase"] = func(variableData interface{}) interface{} {

		return strings.ToLower(variableData.(string))
	}
}
