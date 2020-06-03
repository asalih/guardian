package transformations

func init() {
	TransformationMaps.funcMap["none"] = func(variableData interface{}) interface{} {
		return variableData
	}
}
