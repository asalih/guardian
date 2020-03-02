package transformations

func (transform *TransformationMap) loadCmdLine() {
	transform.funcMap["cmdLine"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
