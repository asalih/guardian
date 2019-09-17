package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/asalih/guardian/request"
)

var dialer = &net.Dialer{
	Timeout:   30 * time.Second,
	KeepAlive: 30 * time.Second,
	DualStack: true,
}

/*GuardianHandler Guardian HTTPS Handler is the transport handler*/
type GuardianHandler struct {
	DB                 *DBHelper
	IsHTTPPortListener bool
}

/*NewGuardianHandler Https Guardian handler init*/
func NewGuardianHandler(isHTTPPortListener bool) *GuardianHandler {
	return &GuardianHandler{&DBHelper{}, isHTTPPortListener}
}

func (h GuardianHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := h.DB.GetTarget(r.Host)

	if target == nil {
		fmt.Fprintf(w, "Your application not authorized yet! Check your implementation. %s", r.URL.Path)

		return
	}

	if target.UseHTTPS && h.IsHTTPPortListener {
		redirectToURI := "https://" + r.Host

		if r.RequestURI != "" {
			redirectToURI += r.RequestURI
		}

		http.Redirect(w, r, redirectToURI, 301)

		return
	}

	requestIsNotSafe := request.NewRequestChecker(w, r).Handle()

	if requestIsNotSafe {
		return
	}

	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {

		uri, ferr := url.Parse(addr)
		if ferr != nil {
			panic(ferr)
		}

		addr = target.OriginIPAddress + ":" + uri.Opaque

		return dialer.DialContext(ctx, network, addr)
	}

	uriToReq := r.Host

	if h.IsHTTPPortListener {
		uriToReq = "http://" + uriToReq
	} else {
		uriToReq = "https://" + uriToReq
	}

	if r.RequestURI != "" {
		uriToReq += r.RequestURI
	}

	transportResponse := h.transportRequest(uriToReq, w, r)

	if transportResponse == nil {
		return
	}

	//TODO: Response check
}

//TransportRequest Transports the incoming request
func (h GuardianHandler) transportRequest(uriToReq string, incomingWriter http.ResponseWriter, incomingRequest *http.Request) *http.Response {
	var response *http.Response
	var err error
	var req *http.Request

	//timeout is 45 secs for to pass to origin server.
	client := &http.Client{
		Timeout: time.Second * 45,
	}

	req, err = http.NewRequest(incomingRequest.Method, uriToReq, incomingRequest.Body)
	for name, value := range incomingRequest.Header {
		//TODO: Do not pass the headers except whitelisted
		req.Header.Set(name, value[0])
	}
	response, err = client.Do(req)
	defer incomingRequest.Body.Close()

	if err != nil {
		http.Error(incomingWriter, err.Error(), http.StatusInternalServerError)
		return nil
	}

	for k, v := range response.Header {
		incomingWriter.Header().Set(k, v[0])
	}
	incomingWriter.WriteHeader(response.StatusCode)
	io.Copy(incomingWriter, response.Body)

	return response
}
