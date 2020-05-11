package transformations

func init() {
	TransformationMaps.funcMap["sqlHexDecode"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
