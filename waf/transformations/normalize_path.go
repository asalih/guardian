package transformations

func init() {
	fn := func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}

	TransformationMaps.funcMap["normalizePath"] = fn
	TransformationMaps.funcMap["normalisePath"] = fn
}
