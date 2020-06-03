package transformations

import "github.com/asalih/guardian/helpers"

func init() {
	TransformationMaps.funcMap["htmlEntityDecode"] = func(variableData interface{}) interface{} {
		return htmlEntitiesDecode(variableData.(string))
	}
}

func htmlEntitiesDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		//If the start of a html encoded entity
		if input[index] == '&' && index+2 < inputLength {
			//If next char is # then we continue parsing
			if input[index+1] == '#' {
				if input[index+2] == 'X' || input[index+2] == 'x' {
					//Try numeric parsing
					left := index + 2
					right := left

					//While we have valid digits move the right pointer right
					//Unless we already have 4 digits
					for right+1 < inputLength && left-right < 4 && helpers.ValidHex(input[right+1]) {
						right++
					}

					//If we have at least one digit we decode
					if right > left {

						code, xv, fact := 0, 0, 1

						for i := 0; left < right-i; i++ {

							//map A-F to 10-16
							//map a-f to 10-16
							//map '1'-'9' to 1-9
							if input[right-i] >= 97 {
								xv = int(input[right-i]) - 97 + 10
							} else if input[right] >= 65 {
								xv = int(input[right-i]) - 65 + 10
							} else {
								xv = int(input[right-i]) - 48
							}

							code += xv * fact
							fact *= 16
						}

						unicodeString := string(code)
						for i := 0; i < len(unicodeString); i++ {
							newString[newIndex] = unicodeString[i]
							newIndex++
						}

						index += 3 + (right - left)

						if index < inputLength && input[index] == ';' {
							index++
						}
					} else {
						//if we have no valid digits we add the & and # to the new string
						newString[newIndex] = input[index]
						newIndex++
						index++

						newString[newIndex] = input[index]
						newIndex++
						index++
					}

				} else {
					//Try numeric parsing
					left := index + 1
					right := left

					//While we have valid digits move the right pointer right
					//Unless we already have 4 digits
					for right+1 < inputLength && left-right < 4 && helpers.IsDigit(input[right+1]) {
						right++
					}

					//If we have at least one digit we decode
					if right > left {

						code, xv, fact := 0, 0, 1

						for i := 0; left < right-i; i++ {

							//map '1'-'9' to 1-9
							//map '1'-'9' to 1-9
							xv = int(input[right-i]) - 48

							code += xv * fact
							fact *= 10
						}

						unicodeString := string(code)
						for i := 0; i < len(unicodeString); i++ {
							newString[newIndex] = unicodeString[i]
							newIndex++
						}

						index += 2 + (right - left)

						if index < inputLength && input[index] == ';' {
							index++
						}
					} else {
						//if we have no valid digits we add the & and # to the new string
						newString[newIndex] = input[index]
						newIndex++
						index++

						newString[newIndex] = input[index]
						newIndex++
						index++
					}
				}
			} else {
				//Try to match to predefined xml entities
				//TODO implement full html predefined entities list
				match := false
				if index+3 < inputLength {
					if input[index+1:index+3] == "gt" {
						newString[newIndex] = '>'
						newIndex++
						index += 3
						match = true
					} else if input[index+1:index+3] == "lt" {
						newString[newIndex] = '<'
						newIndex++
						index += 3
						match = true
					} else if index+4 < inputLength {
						if input[index+1:index+4] == "amp" {
							newString[newIndex] = '&'
							newIndex++
							index += 4
							match = true
						} else if index+5 <= inputLength {
							if input[index+1:index+5] == "quot" {
								newString[newIndex] = '"'
								newIndex++
								index += 5
								match = true
							} else if input[index+1:index+5] == "apos" {
								newString[newIndex] = '\''
								newIndex++
								index += 5
								match = true
							} else if input[index+1:index+5] == "nbsp" {
								newString[newIndex] = '\xa0'
								newIndex++
								index += 5
								match = true
							}
						}

					}
				}

				if match {
					if index < inputLength && input[index] == ';' {
						index++
					}
				} else {
					//Not part of html encoded entity so copy char
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			}
		} else {
			//Otherwise we add the & and we go to the next char
			newString[newIndex] = input[index]
			newIndex++
			index++
		}
	}

	return string(newString[0:newIndex])
}
