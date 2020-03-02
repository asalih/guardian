package transformations

import "html"

func (transform *TransformationMap) loadHTMLEntityDecode() {
	transform.funcMap["htmlEntityDecode"] = func(variableData interface{}) interface{} {
		return html.UnescapeString(variableData.(string))
	}
}
