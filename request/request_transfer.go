package request

import (
	"mime"
	"net/http"

	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/names"
)

//RequestTransfer Request variable object
type RequestTransfer map[string]interface{}

func NewRequestTransfer(request *http.Request) RequestTransfer {
	transferData := RequestTransfer{}

	transferData[names.REQUEST_COOKIES] = request.Cookies()
	transferData[names.REQUEST_URI] = request.RequestURI

	isHTTPS := request.TLS != nil

	transferData[names.REQUEST_URI] = request.RequestURI

	if isHTTPS {
		transferData[names.REQUEST_URI_RAW] = "https://" + request.Host + request.RequestURI
		transferData[names.REQUEST_PROTOCOL] = "https"
	} else {
		transferData[names.REQUEST_URI_RAW] = "http://" + request.Host + request.RequestURI
		transferData[names.REQUEST_PROTOCOL] = "http"
	}

	transferData[names.QUERY_STRING] = request.URL.RawQuery
	transferData[names.REQUEST_BASENAME] = request.URL.Path
	transferData[names.REQUEST_COOKIES_NAMES] = models.GetCookiesNames(request.Cookies())
	transferData[names.REQUEST_HEADERS] = request.Header
	transferData[names.REQUEST_HEADERS_NAMES] = models.GetHeadersNames(request.Header)

	mediaType, mediaParams, err := mime.ParseMediaType(request.Header.Get("ContentType"))
	transferData[names.REQUEST_BODY_TYPE] = mediaType
	transferData[names.MULTIPART_BOUNDARY] = mediaParams["boundary"]

	if err != nil {
		transferData[names.MULTIPART_ERROR] = 1
	} else {
		transferData[names.MULTIPART_ERROR] = 0
	}

	return transferData
}

func (transfer RequestTransfer) Get(key string) string {
	return transfer[key].(string)
}

func (transfer RequestTransfer) GetMap(key string) map[string]string {
	return transfer[key].(map[string]string)
}
