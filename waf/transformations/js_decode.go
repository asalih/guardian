package transformations

func init() {
	TransformationMaps.funcMap["jsDecode"] = func(variableData interface{}) interface{} {
		//Not implemented
		return variableData.(string)
	}
}
