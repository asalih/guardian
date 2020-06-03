package transformations

import (
	"strings"
)

func init() {
	TransformationMaps.funcMap["lowercase"] = func(variableData interface{}) interface{} {
		return strings.ToLower(variableData.(string))
	}
}
