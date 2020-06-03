package transformations

import "crypto/md5"

func init() {
	TransformationMaps.funcMap["md5"] = func(variableData interface{}) interface{} {
		return md5.Sum([]byte(variableData.(string)))
	}
}
