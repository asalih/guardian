package transformations

func (transform *TransformationMap) loadUtf8ToUnicode() {
	transform.funcMap["utf8toUnicode"] = func(variableData interface{}) interface{} {

		//TODO Not implemented
		return variableData
	}
}
