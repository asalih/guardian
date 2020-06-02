package transformations

import "github.com/asalih/guardian/helpers"

func init() {
	TransformationMaps.funcMap["cssDecode"] = func(variableData interface{}) interface{} {
		return cssDecode(variableData.(string))
	}
}

func cssDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		// Is the character a backslash?
		if input[index] == '\\' {

			// Is there at least one more byte?
			if index+1 < inputLength {
				// We are not going to need the backslash.
				index++

				// We are not going to need the backslash.
				j := 0
				for j < 6 && index+j < inputLength && helpers.ValidHex(input[index+j]) {
					j++
				}

				// We have at least one valid hexadecimal character.
				if j > 0 {
					fullcheck := false

					// For now just use the last two bytes
					switch j {
					case 1:
						// Number of hex characters
						newString[newIndex] = helpers.XSingle2c(input[index])
						newIndex++
						break

					case 2:
						//Use the last two from the end
						newString[newIndex] = helpers.X2c(input[index+j-2], input[index+j-1])
						newIndex++
						break
					case 3:
						//Use the last two from the end
						newString[newIndex] = helpers.X2c(input[index+j-2], input[index+j-1])
						newIndex++
						break
					case 4:
						// Use the last two from the end, but request a full width check.
						newString[newIndex] = helpers.X2c(input[index+j-2], input[index+j-1])
						fullcheck = true
						break
					case 5:
						/* Use the last two from the end, but request
						* a full width check if the number is greater
						* or equal to 0xFFFF.
						 */
						newString[newIndex] = helpers.X2c(input[index+j-2], input[index+j-1])

						if input[index] == '0' {
							fullcheck = true
						} else {
							newIndex++
						}

						break
					case 6:
						/* Use the last two from the end, but request
						 * a full width check if the number is greater
						 * or equal to 0xFFFF.
						 */
						newString[newIndex] = helpers.X2c(input[index+j-2], input[index+j-1])
						if input[index] == '0' && input[index+1] == '0' {
							fullcheck = true
						} else {
							newIndex++
						}
					}

					// Full width ASCII (0xff01 - 0xff5e) needs 0x20 added
					if fullcheck {
						if (newString[newIndex] > 0x00) && (newString[newIndex] < 0x5f) &&
							(input[index+j-3] == 'f' || input[index+j-3] == 'F') &&
							(input[index+j-4] == 'f' || input[index+j-4] == 'F') {

							newString[newIndex] += 0x20
						}

						newIndex++
					}

					// We must ignore a single whitespace after a hex escape
					if index+j < inputLength && helpers.IsSpace(input[index+j]) {
						j++
					}

					// Move over.
					index += j

				} else if input[index] == '\n' { // No hexadecimal digits after backslash

					// A newline character following backslash is ignored.
					index++

				} else { // The character after backslash is not a hexadecimal digit, nor a newline.

					// Use one character after backslash as is.
					newString[newIndex] = input[index]
					newIndex++
					index++
				}

			} else { // No characters after backslash.

				// Do not include backslash in output (continuation to nothing)
				index++
			}

		} else { // Character is not a backslash.
			//Copy one normal character to output.
			newString[newIndex] = input[index]
			newIndex++
			index++
		}
	}

	return string(newString[0:newIndex])
}
