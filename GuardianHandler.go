package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
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

	handleResult := NewRequestChecker(w, r).Handle()

	if !handleResult {
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

	resp, gerr := http.Get(uriToReq)

	if gerr != nil {
		//Decide what to do in error case
	}

	defer resp.Body.Close()

	for name, values := range resp.Header {
		w.Header()[name] = values
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}
