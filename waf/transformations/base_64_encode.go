package transformations

import (
	"encoding/base64"
)

func (transform *TransformationMap) loadBase64Encode() {
	transform.funcMap["base64Encode"] = func(variableData interface{}) interface{} {

		return base64.StdEncoding.EncodeToString([]byte(variableData.(string)))

	}
}
