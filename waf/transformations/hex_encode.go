package transformations

import "encoding/hex"

func init() {
	TransformationMaps.funcMap["hexEncode"] = func(variableData interface{}) interface{} {
		switch v := variableData.(type) {
		case string:
			return hex.EncodeToString([]byte(v))

		case []byte:
			return hex.EncodeToString(v)
		}

		return variableData
	}
}
