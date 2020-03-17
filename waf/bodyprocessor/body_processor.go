package bodyprocessor

import "net/http"

//IBodyProcessor Body processor
type IBodyProcessor interface {
	GetBody() map[string][]string
	GetBodyBuffer() []byte
}

//NewBodyProcessor ...
func NewBodyProcessor(r *http.Request) IBodyProcessor {
	if r.Header.Get("Content-Type") == "application/json" {
		return &JSONBodyProcessor{r, nil, nil}
	} else if r.Header.Get("Content-Type") == "multipart/form-data" {
		return &MultipartProcessor{r, nil}
	}

	return &URLEncodedProcessor{r, nil}
}
