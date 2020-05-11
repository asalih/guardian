package transformations

import (
	"bytes"
)

func init() {
	TransformationMaps.funcMap["removeNulls"] = func(variableData interface{}) interface{} {
		return string(bytes.Trim([]byte(variableData.(string)), "\x00"))

	}
}
