package transformations

func (transform *TransformationMap) loadLength() {
	transform.funcMap["length"] = func(variableData interface{}) interface{} {

		return len(variableData.(string))
	}
}
