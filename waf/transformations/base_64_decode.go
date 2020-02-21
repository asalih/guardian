package transformations

import (
	"encoding/base64"
)

func (transform *TransformationMap) loadBase64Decode() {
	transform.funcMap["base64Decode"] = func(variableData interface{}) interface{} {
		str, err := base64.StdEncoding.DecodeString(variableData.(string))
		if err != nil {
			return nil
		}
		return str
	}
}
