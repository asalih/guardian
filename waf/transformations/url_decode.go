package transformations

import (
	"net/url"
)

func (transform *TransformationMap) loadUrlDecode() {
	transform.funcMap["urlDecode"] = func(variableData interface{}) interface{} {

		result, _ := url.QueryUnescape(variableData.(string))

		return result
	}
}
