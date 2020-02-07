package helpers

import (
	"net/http"
	"testing"
	"time"
)

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

func TestUnEscapeRawValue(t *testing.T) {
	testDataMap := map[string]map[string]interface{}{
		"MATCH1": {"pattern": "%%", "expect": "%%"},
		"MATCH2": {"pattern": "%'", "expect": "%"},
		"MATCH3": {"pattern": `%"`, "expect": `%`},
		"EMPTY":  {"pattern": "", "expect": ""},
		"SPACE":  {"pattern": " ", "expect": " "},
		"PLUS":   {"pattern": "+", "expect": " "},
	}

	for _, testData := range testDataMap {

		result := UnEscapeRawValue(testData["pattern"].(string))

		if result == testData["expect"].(string) {
			t.Logf("UnEscapeRawValue(\"%v\") PASSED, expected %v got %v.",
				testData["pattern"], testData["expect"], result)
		} else {
			t.Errorf("UnEscapeRawValue(\"%v\") FAILED, expected %v got %v.",
				testData["pattern"], testData["expect"], result)
		}
	}
}

func TestPreProcessString(t *testing.T) {
	testDataMap := map[string]map[string]interface{}{
		"MATCH1": {"pattern": `'`, "expect": ``},
		"MATCH2": {"pattern": `"`, "expect": ``},
		"MATCH3": {"pattern": `/**/`, "expect": ` `},
		"MATCH4": {"pattern": `+`, "expect": ` `},
		"EMPTY":  {"pattern": "", "expect": ""},
		"SPACE":  {"pattern": " ", "expect": " "},
	}

	for _, testData := range testDataMap {

		result := PreProcessString(testData["pattern"].(string))

		if result == testData["expect"].(string) {
			t.Logf("PreProcessString(\"%v\") PASSED, expected %v got %v.",
				testData["pattern"], testData["expect"], result)
		} else {
			t.Errorf("PreProcessString(\"%v\") FAILED, expected %v got %v.",
				testData["pattern"], testData["expect"], result)
		}
	}
}

func TestHeadersToString(t *testing.T) {
	req, _ := http.NewRequest("GET", "www.netsparker.com", nil)

	req.Header.Add("Content-Type", "application/json")

	expect := "Content-Type: application/json "

	result := HeadersToString(req.Header)

	if result == expect {
		t.Logf("HeadersToString(\"%v\") PASSED, expected %v got %v.",
			req.Header, expect, result)
	} else {
		t.Errorf("PreProcessString(\"%v\") FAILED, expected %v got %v.",
			req.Header, expect, result)
	}
}

func TestCookiesToString(t *testing.T) {
	req, _ := http.NewRequest("GET", "www.netsparker.com", nil)

	req.AddCookie(&http.Cookie{Name: "cookie1", Value: "c1"})
	req.AddCookie(&http.Cookie{Name: "cookie2", Value: "c2"})

	expect := "cookie1=c1 cookie2=c2 "

	result := CookiesToString(req.Cookies())

	if result == expect {
		t.Logf("CookiesToString(\"%v\") PASSED, expected %v got %v.",
			req.Cookies(), expect, result)
	} else {
		t.Errorf("CookiesToString(\"%v\") FAILED, expected %v got %v.",
			req.Cookies(), expect, result)
	}
}

func TestCalcTime(t *testing.T) {
	start := time.Now()

	time.Sleep(1 * time.Second)

	end := time.Now()

	result := CalcTime(start, end)
	expect := int64(1000)

	if result >= expect {
		t.Logf("CalcTime(\"%v\") PASSED, expected %v got %v.",
			start, expect, result)
	} else {
		t.Errorf("CalcTime(\"%v\") FAILED, expected %v got %v.",
			start, expect, result)
	}
}

func TestCalcTimeNow(t *testing.T) {
	ti := time.Now()

	time.Sleep(1 * time.Second)

	result := CalcTimeNow(ti)
	expect := int64(1000)

	if result >= expect {
		t.Logf("CalcTime(\"%v\") PASSED, expected %v got %v.",
			ti, expect, result)
	} else {
		t.Errorf("CalcTime(\"%v\") FAILED, expected %v got %v.",
			ti, expect, result)
	}
}
