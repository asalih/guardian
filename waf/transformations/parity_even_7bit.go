package transformations

func (transform *TransformationMap) loadParityEven7bit() {
	transform.funcMap["parityEven7bit"] = func(variableData interface{}) interface{} {

		input := []byte(variableData.(string))
		inputLen := len(input)

		for i := 0; i < inputLen; i++ {
			x := input[i]

			input[i] ^= input[i] >> 4
			input[i] &= 0xf

			if (0x6996>>input[i])&1 > 0 {
				input[i] = x | 0x80
			} else {
				input[i] = x & 0x7f
			}
		}
		return string(input)
	}
}
