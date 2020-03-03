package transformations

import (
	"crypto/sha1"
)

func (transform *TransformationMap) loadSHA1() {
	transform.funcMap["sha1"] = func(variableData interface{}) interface{} {

		return sha1.Sum([]byte(variableData.(string)))
	}
}
