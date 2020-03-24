package bodyprocessor

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//ResponseBodyProcessor Response body processor
type ResponseBodyProcessor struct {
	response     *http.Response
	bodyBuffer   []byte
	hasBodyError bool
}

//GetBody ...
func (p *ResponseBodyProcessor) GetBody() map[string][]string {

	//Not implemented
	return nil
}

//GetBodyBuffer ...
func (p *ResponseBodyProcessor) GetBodyBuffer() []byte {

	if len(p.bodyBuffer) > 0 {
		return p.bodyBuffer
	}

	bodyBytes, _ := ioutil.ReadAll(p.response.Body)
	p.response.Body.Close() //  must close
	p.bodyBuffer = bodyBytes
	p.response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return p.bodyBuffer
}

//HasBodyError ...
func (p *ResponseBodyProcessor) HasBodyError() bool {
	return p.hasBodyError
}
