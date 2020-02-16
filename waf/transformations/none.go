package transformations

func (transform *TransformationMap) loadNone() {
	transform.funcMap["none"] = func(variableData interface{}) interface{} {

		return variableData
	}
}
