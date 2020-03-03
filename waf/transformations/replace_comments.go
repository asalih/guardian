package transformations

import (
	"regexp"
)

var re = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")

func (transform *TransformationMap) loadReplaceComments() {
	transform.funcMap["replaceComments"] = func(variableData interface{}) interface{} {

		str := variableData.(string)
		return re.ReplaceAllString(str, " ")

	}
}
