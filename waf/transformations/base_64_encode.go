package transformations

import (
	"encoding/base64"
)

func init() {
	TransformationMaps.funcMap["base64Encode"] = func(variableData interface{}) interface{} {

		return base64.StdEncoding.EncodeToString([]byte(variableData.(string)))

	}
}
