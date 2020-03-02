package transformations

func (transform *TransformationMap) loadEscapeSeqDecode() {
	transform.funcMap["escapeSeqDecode"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
