package transformations

import (
	"github.com/asalih/guardian/helpers"
)

func init() {
	TransformationMaps.funcMap["urlDecodeUni"] = func(variableData interface{}) interface{} {
		return urlDecodeUni(variableData.(string), 20127)
	}
}

func urlDecodeUni(
	input string,
	unicodeCodePage int,
) string {

	if input == "" {
		return ""
	}

	newString := make([]byte, len(input))

	index, newIndex, xv, code, fact := 0, 0, 0, 0, 0
	hmap := byte(0)
	hmapFound := false

	inputLength := len(input)

	for index < inputLength {
		if input[index] == '%' {
			// Character is a percent sign.

			if (index+1) < inputLength && ((input[index+1] == 'u') || (input[index+1] == 'U')) {
				// IIS-specific %u encoding.

				if index+5 < inputLength {
					// We have at least 4 data bytes.
					if helpers.ValidHex(input[index+2]) &&
						helpers.ValidHex(input[index+3]) &&
						helpers.ValidHex(input[index+4]) &&
						helpers.ValidHex(input[index+5]) {

						code = 0
						fact = 1
						hmapFound = false

						if len(unicodemap) > 0 && unicodeCodePage > 0 {

							for i := 5; i >= 2; i-- {
								if helpers.ValidHex(input[index+i]) {
									if input[index+i] >= 97 {
										xv = int(input[index+i]) - 97 + 10
									} else if input[index+i] >= 65 {
										xv = int(input[index+i]) - 65 + 10
									} else {
										xv = int(input[index+i]) - 48
									}
									code += xv * fact
									fact *= 16
								}
							}

							if code >= 0 && code <= 65535 {
								hmap, hmapFound = unicodemap[unicodeCodePage][code]
							}
						}

						if hmapFound {
							newString[newIndex] = hmap
						} else {
							// We first make use of the lower byte here, ignoring the higher byte.
							newString[newIndex] = helpers.X2c(input[index+4], input[index+5])

							// Full width ASCII (ff01 - ff5e) needs 0x20 added
							if (newString[newIndex] > 0x00) &&
								(newString[newIndex] < 0x5f) &&
								(input[index+2] == 'f' || input[index+2] == 'F' &&
									(input[index+3] == 'f' || input[index+3] == 'F')) {
								newString[newIndex] += 0x20

							}
						}
						newIndex++
						index += 6
					} else {
						// Invalid data, skip %u.
						newString[newIndex] = input[index]
						newIndex++
						newString[newIndex] = input[index+1]
						newIndex++
						index += 2
					}
				} else {
					// Not enough bytes (4 data bytes), skip %u.
					newString[newIndex] = input[index]
					newIndex++
					newString[newIndex] = input[index+1]
					newIndex++
					index += 2
				}
			} else {
				// Standard URL encoding.

				// Are there enough bytes available?
				if index+2 < inputLength {
					// Yes

					// Decode a %xx combo only if it is valid.
					c1, c2 := input[index+1], input[index+2]

					if helpers.ValidHex(c1) && helpers.ValidHex(c2) {
						newString[newIndex] = helpers.X2c(c1, c2)
						newIndex++
						index += 3
					} else {
						// Not a valid encoding, skip this %
						newString[newIndex] = input[index]
						newIndex++
						index++
					}
				} else {
					// Not enough bytes available, skip this %
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			}
		} else {

			// Character is not a percent sign.
			if input[index] == '+' {
				newString[newIndex] = ' '
			} else {
				newString[newIndex] = input[index]
			}
			newIndex++
			index++
		}
	}

	return string(newString[:newIndex])
}
