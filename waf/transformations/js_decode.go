package transformations

import "github.com/asalih/guardian/helpers"

func init() {
	TransformationMaps.funcMap["jsDecode"] = func(variableData interface{}) interface{} {
		//Not implemented
		return jsDecode(variableData.(string))
	}
}

func jsDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		if input[index] == '\\' {

			/* \uHHHH unicode escape sequence*/
			if index+5 < inputLength && input[index+1] == 'u' &&
				helpers.ValidHex(input[index+2]) && helpers.ValidHex(input[index+3]) &&
				helpers.ValidHex(input[index+4]) && helpers.ValidHex(input[index+5]) {

				/* Use only the lower byte. */
				newString[newIndex] = helpers.X2c(input[index+4], input[index+5])

				/* Full width ASCII (ff01 - ff5e) needs 0x20 added */
				if newString[newIndex] > 0x00 && newString[newIndex] < 0x5f &&
					(input[index+2] == 'f' || input[index+2] == 'F') &&
					(input[index+3] == 'f' || input[index+3] == 'F') {

					newString[newIndex] += 0x20
				}

				newIndex++
				index += 6
				/* \xHH hex secapte sequence*/
			} else if index+3 < inputLength && input[index+1] == 'x' &&
				helpers.ValidHex(input[index+2]) && helpers.ValidHex(input[index+3]) {

				newString[newIndex] = helpers.X2c(input[index+2], input[index+3])
				newIndex++
				index += 4

				/* \OOO (only one byte, \000 - \377) */
			} else if index+1 < inputLength && helpers.IsODidit(input[index+1]) {

				//TODO check this alloc, it is probably not necessary
				buf := make([]byte, 4)
				j := 0
				for index+1+j < inputLength && j < 3 {
					buf[j] = input[index+1+j]
					j++
					if index+1+j < inputLength && !helpers.IsODidit(input[index+1+j]) {
						break
					}
				}

				if j > 0 {
					/* Do not use 3 characters if we will be > 1 byte */
					if (j == 3) && (buf[0] > '3') {
						j = 2
						buf[j] = '\x00'
					}

					code, xv, fact := 0, 0, 1

					for i := 1; i <= j; i++ {

						//map '1'-'9' to 1-9
						xv = int(buf[j-i]) - 48

						code += xv * fact
						fact *= 8
					}

					newString[newIndex] = byte(code)
					newIndex++
					index += 1 + j
				}

				/* \C */
			} else if index+1 < inputLength {
				c := input[index+1]
				switch input[index+1] {
				case 'a':
					c = '\a'
					break
				case 'b':
					c = '\b'
					break
				case 'f':
					c = '\f'
					break
				case 'n':
					c = '\n'
					break
				case 'r':
					c = '\r'
					break
				case 't':
					c = '\t'
					break
				case 'v':
					c = '\v'
					break
					/* The remaining (\?,\\,\',\") are just a removal
					 * of the escape char which is default.
					 */
				}

				newString[newIndex] = c
				newIndex++
				index += 2
			} else {
				/* Not enough bytes */
				for index < inputLength {
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			}
		} else {
			newString[newIndex] = input[index]
			newIndex++
			index++
		}
	}

	return string(newString[0:newIndex])
}
