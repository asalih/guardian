package transformations

import (
	"strings"
)

func init() {
	TransformationMaps.funcMap["removeCommentsChar"] = func(variableData interface{}) interface{} {
		str := variableData.(string)
		replacer := strings.NewReplacer("/*", "", "*/", "", "--", "", "#", "")

		return replacer.Replace(str)
	}
}
