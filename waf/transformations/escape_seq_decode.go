package transformations

func init() {
	TransformationMaps.funcMap["escapeSeqDecode"] = func(variableData interface{}) interface{} {
		//TODO Not implemented
		return variableData.(string)
	}
}
