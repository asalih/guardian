package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*RequestCheck Cheks the requests init*/
type RequestCheck struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

/*NewRequestChecker Request checker initializer*/
func NewRequestChecker(w http.ResponseWriter, r *http.Request) *RequestCheck {
	return &RequestCheck{w, r}
}

/*Handle Request checker handler func*/
func (r RequestCheck) Handle() bool {

	notSafeForScript := r.lookRequest()

	if notSafeForScript {
		fmt.Fprintf(r.ResponseWriter, "Not nice. %s", r.Request.URL.Path)

		return false
	}

	return true
}

func (r RequestCheck) lookRequest() bool {
	return strings.Contains(r.Request.URL.RawQuery, "%3Cscript%3E")
}
