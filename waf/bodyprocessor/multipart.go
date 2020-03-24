package bodyprocessor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//MultipartProcessor URL Encoded body parser
type MultipartProcessor struct {
	request      *http.Request
	bodyBuffer   []byte
	hasBodyError bool
}

//GetBody ...
func (p *MultipartProcessor) GetBody() map[string][]string {

	if p.request.Form != nil && p.request.PostForm != nil {
		return nil
	}

	p.GetBodyBuffer()
	err := p.request.ParseMultipartForm(1024 * 1024 * 4)

	if err != nil {
		fmt.Println(err)

		p.hasBodyError = true
	}

	return map[string][]string(p.request.Form)
}

//GetBodyBuffer ...
func (p *MultipartProcessor) GetBodyBuffer() []byte {

	if len(p.bodyBuffer) > 0 {
		return p.bodyBuffer
	}

	bodyBytes, _ := ioutil.ReadAll(p.request.Body)
	p.request.Body.Close() //  must close
	p.bodyBuffer = bodyBytes
	p.request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return p.bodyBuffer
}

//HasBodyError ...
func (p *MultipartProcessor) HasBodyError() bool {
	return p.hasBodyError
}
