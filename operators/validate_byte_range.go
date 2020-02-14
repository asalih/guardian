package operators

import (
	"strconv"
	"strings"
)

func (opMap *OperatorMap) loadValidateByteRange() {
	opMap.funcMap["validateByteRange"] = func(expression interface{}, variableData interface{}) bool {

		rangeMap := getRange(expression.(string))
		data := []byte(variableData.(string))

		for i := 0; i < len(data); i++ {

			if rangeMap[data[i]] == 0 {
				return true
			}
		}

		return false
	}
}

func getRange(rangeExpression string) map[byte]byte {
	items := strings.Split(rangeExpression, ",")

	maps := make(map[byte]byte)
	for i := 0; i < len(items); i++ {
		ritem := items[i]

		if strings.Contains(ritem, "-") {
			ranges := strings.Split(ritem, "-")

			start, _ := strconv.Atoi(strings.Trim(ranges[0], " "))
			end, _ := strconv.Atoi(strings.Trim(ranges[1], " "))

			for j := start; j <= end; j++ {
				maps[byte(j)] = 1
			}

		} else {
			ibyte, _ := strconv.Atoi(ritem)
			maps[byte(ibyte)] = 1
		}
	}

	return maps
}
