package operators

func init() {
	fn := func(expression interface{}, variableData interface{}) bool {

		fileCache := DataFileCaches[expression.(string)]

		if fileCache == nil || fileCache.Matcher == nil {
			return false
		}

		hits := fileCache.Matcher.Match([]byte(variableData.(string)))

		if len(hits) > 0 {
			return true
		}

		return false
	}

	OperatorMaps.funcMap["pmf"] = fn
	OperatorMaps.funcMap["pmFromFile"] = fn
}
