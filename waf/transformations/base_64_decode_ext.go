package transformations

func (transform *TransformationMap) loadBase64DecodeExt() {
	transform.funcMap["base64DecodeExt"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
