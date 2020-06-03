package transformations

import (
	"strings"
	"unicode"
)

func init() {
	TransformationMaps.funcMap["removeWhitespace"] = func(variableData interface{}) interface{} {
		str := variableData.(string)
		var b strings.Builder
		b.Grow(len(str))

		for _, ch := range str {
			if !unicode.IsSpace(ch) {
				b.WriteRune(ch)
			}
		}

		return b.String()
	}
}
