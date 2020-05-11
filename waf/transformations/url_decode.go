package transformations

import (
	"net/url"
)

func init() {
	TransformationMaps.funcMap["urlDecode"] = func(variableData interface{}) interface{} {

		result, _ := url.QueryUnescape(variableData.(string))

		return result
	}
}
