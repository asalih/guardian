package models

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

//IsMatch checks for matching string
func IsMatch(pattern string, str string) (bool, error) {

	matched, _ := regexp.MatchString(pattern, str)
	fmt.Printf("%s || %s || %s\n", pattern, str, matched)
	return matched, nil
}

//UnEscapeRawValue unescapes raw query
func UnEscapeRawValue(rawQuery string) string {
	rawQuery = strings.Replace(rawQuery, "%%", "%25%", -1)
	rawQuery = strings.Replace(rawQuery, "%'", "%25'", -1)
	rawQuery = strings.Replace(rawQuery, `%"`, `%25"`, -1)
	re := regexp.MustCompile(`%$`)
	rawQuery = re.ReplaceAllString(rawQuery, `%25`)
	decodeQuery, _ := url.QueryUnescape(rawQuery)

	decodeQuery = PreProcessString(decodeQuery)
	//fmt.Println("UnEscapeRawValue decodeQuery", decodeQuery)
	return decodeQuery
}

// PreProcessString ...
func PreProcessString(value string) string {
	value2 := strings.Replace(value, `'`, ``, -1)
	value2 = strings.Replace(value2, `"`, ``, -1)
	value2 = strings.Replace(value2, `+`, ` `, -1)
	value2 = strings.Replace(value2, `/**/`, ` `, -1)
	return value2
}

//HeadersToString ...
func HeadersToString(header http.Header) (res string) {
	for name, values := range header {
		for _, value := range values {
			res += fmt.Sprintf("%s: %s", name, value)
		}
	}
	return
}

//CookiesToString ...
func CookiesToString(cookie []*http.Cookie) (res string) {
	for _, values := range cookie {
		res += fmt.Sprintf("%s=%s ", values.Name, values.Value)
	}
	return
}

//CalcTime ...
func CalcTime(start time.Time) int64 {
	return time.Since(start).Nanoseconds() / int64(time.Millisecond)
}
