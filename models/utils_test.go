package models

import "testing"

func TestIsMatchWithValues(t *testing.T) {
	testDataMap := map[string]map[string]interface{}{
		"EMPTY TEST":     {"pattern": "", "str": "", "expect": true},
		"NOT MATCH TEST": {"pattern": "ahmet", "str": "salih", "expect": false},
		"MATCH TEST":     {"pattern": ".*", "str": "netsparker", "expect": true},
	}

	for _, testData := range testDataMap {

		result, err := IsMatch(testData["pattern"].(string), testData["str"].(string))

		if err != nil {
			t.Errorf("IsMatch(\"%v\", \"%v\") FAILED, expected %v got error %v.",
				testData["pattern"], testData["str"], testData["expect"], err)
		}

		if result == testData["expect"].(bool) {
			t.Logf("IsMatch(\"%v\", \"%v\") PASSED, expected %v got %v.",
				testData["pattern"], testData["str"], testData["expect"], result)
		} else {
			t.Errorf("IsMatch(\"%v\", \"%v\") FAILED, expected %v got %v.",
				testData["pattern"], testData["str"], testData["expect"], result)
		}
	}

}

/*func TestUnEscapeRawValue(t *testing.T) {
	testData := make(map[string]string)
	testData["?q=1&a=<script>alert(1)</script>"] = ""

	result := UnEscapeRawValue("?q=1&a=<script>alert(1)</script>")

	if result {
		t.Log("UnEscapeRawValue(\"\", \"\") PASSED, expected true got true.")
	} else {
		t.Error("IsMatch(\"\", \"\") FAILED, expected true got true.")
	}
}
*/
