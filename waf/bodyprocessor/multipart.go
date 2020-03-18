package bodyprocessor

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//MultipartProcessor URL Encoded body parser
type MultipartProcessor struct {
	request    *http.Request
	bodyBuffer []byte
}

//GetBody ...
func (p *MultipartProcessor) GetBody() map[string][]string {

	if p.request.Form != nil && p.request.PostForm != nil {
		return nil
	}

	p.GetBodyBuffer()
	err := p.request.ParseMultipartForm(1024 * 1024 * 4)

	if err != nil {
		panic(err)
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
