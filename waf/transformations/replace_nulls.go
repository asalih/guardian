package transformations

import (
	"bytes"
)

func init() {
	TransformationMaps.funcMap["replaceNulls"] = func(variableData interface{}) interface{} {
		return string(bytes.ReplaceAll([]byte(variableData.(string)), []byte("\x00"), []byte(" ")))
	}
}
