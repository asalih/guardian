package transformations

import "encoding/hex"

func (transform *TransformationMap) loadHexDecode() {
	transform.funcMap["hexDecode"] = func(variableData interface{}) interface{} {
		decoded, err := hex.DecodeString(variableData.(string))
		if err != nil {
			return variableData
		}

		return decoded
	}
}
