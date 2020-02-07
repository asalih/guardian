package operators

import (
	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadUrlEncoding() {
	opMap.funcMap["validateUrlEncoding"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)

		data := variableData.(string)

		if data == "" {
			return matchResult
		}

		if validateUrlEncoding(variableData.(string)) == 1 {
			matchResult.SetMatch(true)
		}

		return matchResult
	}
}

func validateUrlEncoding(input string) int {
	var i int

	input_length := len(input)

	if (input == "") || (input_length <= 0) {
		return -1
	}

	i = 0
	for i < input_length {
		if input[i] == '%' {
			if i+2 >= input_length {

				/* Not enough bytes. */
				return -3
			} else {
				/* Here we only decode a %xx combination if it is valid,
				 * leaving it as is otherwise.
				 */
				c1 := input[i+1]
				c2 := input[i+2]

				if (((c1 >= '0') && (c1 <= '9')) ||
					((c1 >= 'a') && (c1 <= 'f')) ||
					((c1 >= 'A') && (c1 <= 'F'))) &&
					(((c2 >= '0') && (c2 <= '9')) ||
						((c2 >= 'a') && (c2 <= 'f')) ||
						((c2 >= 'A') && (c2 <= 'F'))) {
					i += 3
				} else {
					/* Non-hexadecimal characters used in encoding. */
					return -2
				}
			}
		} else {
			i++
		}
	}
	return 1
}
