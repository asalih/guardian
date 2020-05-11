package transformations

func init() {
	TransformationMaps.funcMap["base64DecodeExt"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
