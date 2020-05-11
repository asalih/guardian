package transformations

import (
	"crypto/sha1"
)

func init() {
	TransformationMaps.funcMap["sha1"] = func(variableData interface{}) interface{} {

		return sha1.Sum([]byte(variableData.(string)))
	}
}
