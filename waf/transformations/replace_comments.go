package transformations

import (
	"regexp"
)

var re = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")

func init() {
	TransformationMaps.funcMap["replaceComments"] = func(variableData interface{}) interface{} {
		str := variableData.(string)

		return re.ReplaceAllString(str, " ")
	}
}
