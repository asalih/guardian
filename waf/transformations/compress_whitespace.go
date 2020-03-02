package transformations

import (
	"strings"
	"unicode"
)

func (transform *TransformationMap) loadCompressWhitespace() {
	transform.funcMap["compressWhitespace"] = func(variableData interface{}) interface{} {

		str := variableData.(string)
		var b strings.Builder
		b.Grow(len(str))
		inSpace := false
		for _, ch := range str {
			if unicode.IsSpace(ch) {
				if inSpace {
					continue
				}

				inSpace = true
				b.WriteString(" ")
			} else {
				inSpace = false
				b.WriteRune(ch)
			}
		}

		return b.String()
	}
}
