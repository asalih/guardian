package transformations

import (
	"bytes"
)

func (transform *TransformationMap) loadRemoveNulls() {
	transform.funcMap["removeNulls"] = func(variableData interface{}) interface{} {
		return string(bytes.Trim([]byte(variableData.(string)), "\x00"))

	}
}
