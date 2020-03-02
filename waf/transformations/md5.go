package transformations

import "crypto/md5"

func (transform *TransformationMap) loadMD5() {
	transform.funcMap["md5"] = func(variableData interface{}) interface{} {

		return md5.Sum([]byte(variableData.(string)))
	}
}
