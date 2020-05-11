package operators

func init() {
	OperatorMaps.funcMap["validateUrlEncoding"] = func(expression interface{}, variableData interface{}) bool {

		data := variableData.(string)

		if data == "" {
			return false
		}

		if validateURLEncoding(variableData.(string)) == 1 {
			return true
		}

		return false
	}
}

func validateURLEncoding(input string) int {
	var i int

	inputLength := len(input)

	if (input == "") || (inputLength <= 0) {
		return -1
	}

	i = 0
	for i < inputLength {
		if input[i] == '%' {
			if i+2 >= inputLength {

				/* Not enough bytes. */
				return -3
			}

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

		} else {
			i++
		}
	}
	return 1
}
