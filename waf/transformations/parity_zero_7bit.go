package transformations

func init() {
	TransformationMaps.funcMap["parityZero7bit"] = func(variableData interface{}) interface{} {

		input := []byte(variableData.(string))
		inputLen := len(input)

		for i := 0; i < inputLen; i++ {
			input[i] &= 0x7f
		}
		return string(input)
	}
}
