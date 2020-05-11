package transformations

func init() {
	TransformationMaps.funcMap["length"] = func(variableData interface{}) interface{} {

		return len(variableData.(string))
	}
}
