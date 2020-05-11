package transformations

import (
	"encoding/base64"
)

func init() {
	TransformationMaps.funcMap["base64Decode"] = func(variableData interface{}) interface{} {
		str, err := base64.StdEncoding.DecodeString(variableData.(string))
		if err != nil {
			return nil
		}
		return str
	}
}
