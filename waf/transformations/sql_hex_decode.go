package transformations

func (transform *TransformationMap) loadSQLHexDecode() {
	transform.funcMap["sqlHexDecode"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
