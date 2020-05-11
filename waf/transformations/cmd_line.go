package transformations

func init() {
	TransformationMaps.funcMap["cmdLine"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
