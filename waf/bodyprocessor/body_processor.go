package bodyprocessor

import "net/http"

//IBodyProcessor Body processor
type IBodyProcessor interface {
	GetBody() map[string][]string
	HasBodyError() bool
	GetBodyBuffer() []byte
}

//NewBodyProcessor Initiates a body processor by content-type
func NewBodyProcessor(r *http.Request) IBodyProcessor {
	if r.Header.Get("Content-Type") == "application/json" {
		return &JSONBodyProcessor{r, nil, make(map[string][]string), false}
	} else if r.Header.Get("Content-Type") == "multipart/form-data" {
		return &MultipartProcessor{r, nil, false}
	} else if r.Header.Get("Content-Type") == "application/xml" {
		return &XMLBodyProcessor{r, nil, false, nil}
	}

	return &URLEncodedProcessor{r, nil, false}
}

//NewResponseBodyProcessor inits response body processor
func NewResponseBodyProcessor(r *http.Response) IBodyProcessor {
	return &ResponseBodyProcessor{r, nil, false}
}
