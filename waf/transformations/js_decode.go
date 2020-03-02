package transformations

func (transform *TransformationMap) loadJSDecode() {
	transform.funcMap["jsDecode"] = func(variableData interface{}) interface{} {
		//Not implemented
		return variableData.(string)
	}
}
