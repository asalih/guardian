package transformations

import (
	"net/url"
)

func init() {
	TransformationMaps.funcMap["urlDecodeUni"] = func(variableData interface{}) interface{} {

		//TODO: url decode uni has to be reviewed
		result, _ := url.QueryUnescape(variableData.(string))

		return result
	}
}
