package transformations

import (
	"strings"
)

func (transform *TransformationMap) loadRemoveCommentsChar() {
	transform.funcMap["removeCommentsChar"] = func(variableData interface{}) interface{} {
		str := variableData.(string)
		replacer := strings.NewReplacer("/*", "", "*/", "", "--", "", "#", "")

		return replacer.Replace(str)
	}
}
