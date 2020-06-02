package helpers

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
	return regexp.MatchString(pattern, str)
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
			res += fmt.Sprintf("%s: %s ", name, value)
		}
	}
	return
}

//GetHeadersNames Gets the header name
func GetHeadersNames(header http.Header) (res []string) {
	for name := range header {
		res = append(res, name)
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

//GetCookiesNames ...
func GetCookiesNames(cookie []*http.Cookie) (res []string) {
	for _, values := range cookie {
		res = append(res, values.Name)
	}
	return
}

//CalcTime ...
func CalcTime(start time.Time, end time.Time) int64 {
	return end.Sub(start).Nanoseconds() / int64(time.Millisecond)
}

//CalcTimeNow ...
func CalcTimeNow(end time.Time) int64 {
	return time.Since(end).Nanoseconds() / int64(time.Millisecond)
}

//StringContains searches given string in a string slice
func StringContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

//X2c Combines two hex chars into a single byte
func X2c(c1 byte, c2 byte) byte {
	var char byte
	if c1 >= 'A' {
		char = ((c1 & 0xdf) - 'A') + 10
	} else {
		char = c1 - '0'
	}
	char *= 16
	if c2 >= 'A' {
		char += ((c2 & 0xdf) - 'A') + 10
	} else {
		char += c2 - '0'
	}
	return char
}

//XSingle2c Converts a hex char into a byte
func XSingle2c(c byte) byte {
	if c >= 'A' {
		return ((c & 0xdf) - 'A') + 10
	}

	return c - '0'
}

//ValidHex ...
func ValidHex(X byte) bool {
	return ((X >= '0') && (X <= '9')) || ((X >= 'a') && (X <= 'f')) || ((X >= 'A') && (X <= 'F'))
}

//IsDigit Determines given byte is digit
func IsDigit(X byte) bool {
	return (X >= '0') && (X <= '9')
}

//IsODidit checks if the byte is a octo decimal digit
func IsODidit(X byte) bool {
	return (X >= '0') && (X <= '7')
}

/*IsSpace checks for white-space  characters.   In  the  "C"  and  "POSIX"
locales,  these  are:  space,  form-feed ('\f'), newline ('\n'),
carriage return ('\r'), horizontal tab ('\t'), and vertical  tab
('\v'). */
func IsSpace(X byte) bool {
	return X == ' ' || X == '\n' ||
		X == '\r' || X == '\t' ||
		X == '\f' || X == '\v'
}
