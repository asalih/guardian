package transformations

import "html"

func init() {
	TransformationMaps.funcMap["htmlEntityDecode"] = func(variableData interface{}) interface{} {
		return html.UnescapeString(variableData.(string))
	}
}
