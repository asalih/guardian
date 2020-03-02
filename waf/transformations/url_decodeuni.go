package transformations

import (
	"net/url"
)

func (transform *TransformationMap) loadURLDecodeUni() {
	transform.funcMap["urlDecodeUni"] = func(variableData interface{}) interface{} {

		//TODO: url decode uni has to be reviewed
		result, _ := url.QueryUnescape(variableData.(string))

		return result
	}
}
