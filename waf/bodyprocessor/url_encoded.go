package bodyprocessor

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//URLEncodedProcessor URL Encoded body parser
type URLEncodedProcessor struct {
	request    *http.Request
	bodyBuffer []byte
}

//GetBody ...
func (p *URLEncodedProcessor) GetBody() map[string][]string {

	if p.request.Form != nil && p.request.PostForm != nil {
		return map[string][]string(p.request.Form)
	}

	p.GetBodyBuffer()
	err := p.request.ParseForm()

	if err != nil {
		panic(err)
	}

	return map[string][]string(p.request.Form)
}

//GetBodyBuffer ...
func (p *URLEncodedProcessor) GetBodyBuffer() []byte {

	if len(p.bodyBuffer) > 0 {
		return p.bodyBuffer
	}

	bodyBytes, _ := ioutil.ReadAll(p.request.Body)
	p.request.Body.Close() //  must close
	p.bodyBuffer = bodyBytes
	p.request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return p.bodyBuffer
}
