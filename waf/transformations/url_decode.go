package transformations

import "github.com/asalih/guardian/helpers"

func init() {
	TransformationMaps.funcMap["urlDecode"] = func(variableData interface{}) interface{} {
		return urlDecode(variableData.(string))
	}
}

func urlDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		if input[index] == '%' {
			/* Character is a percent sign. */

			/* Are there enough bytes available? */
			if index+2 < inputLength {
				c1 := input[index+1]
				c2 := input[index+2]

				if helpers.ValidHex(c1) && helpers.ValidHex(c2) {
					/* Valid encoding - decode it. */
					newString[newIndex] = helpers.X2c(input[index+1], input[index+2])
					newIndex++
					index += 3
				} else {
					/* Not a valid encoding, skip this % */
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			} else {
				/* Not enough bytes available, copy the raw bytes. */
				newString[newIndex] = input[index]
				newIndex++
				index++
			}
		} else {
			/* Character is not a percent sign. */
			if input[index] == '+' {
				newString[newIndex] = ' '
				newIndex++
			} else {
				newString[newIndex] = input[index]
				newIndex++
			}
			index++
		}
	}

	return string(newString[0:newIndex])
}
