package transformations

func (transform *TransformationMap) loadNormalizePath() {
	fn := func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}

	transform.funcMap["normalizePath"] = fn
	transform.funcMap["normalisePath"] = fn
}
