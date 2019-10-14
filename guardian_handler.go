package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/asalih/guardian/response"

	"github.com/asalih/guardian/data"
	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/request"
)

var dialer = &net.Dialer{
	Timeout:   30 * time.Second,
	KeepAlive: 30 * time.Second,
	DualStack: true,
}

/*GuardianHandler Guardian HTTPS Handler is the transport handler*/
type GuardianHandler struct {
	DB                 *data.DBHelper
	IsHTTPPortListener bool
	CertManager        *autocert.Manager
}

/*NewGuardianHandler Https Guardian handler init*/
func NewGuardianHandler(isHTTPPortListener bool, certManager *autocert.Manager) *GuardianHandler {
	return &GuardianHandler{&data.DBHelper{}, isHTTPPortListener, certManager}
}

func (h GuardianHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Guardian Handler Executing: " + r.Host)

	target := h.DB.GetTarget(r.Host)

	if target == nil {
		fmt.Fprintf(w, "Your application not authorized yet! Check your implementation. %s", r.URL.Path)
		fmt.Println("Unauthorized Application requested." + r.Host)

		return
	}

	if target.AutoCert && h.IsHTTPPortListener {
		fmt.Println("AutoCert in progress. " + r.Host + r.URL.Path)
		h.CertManager.HTTPHandler(nil).ServeHTTP(w, r)
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

	httpLog := models.NewHTTPLog()

	requestIsNotSafe := request.NewRequestChecker(w, r, target).Handle()

	httpLog = httpLog.RequestRulesExecutionEnd()

	if requestIsNotSafe {
		go h.logHTTPRequest(httpLog.Build(target, r, nil))

		return
	}

	uriToReq := r.Host

	if target.Proto == 0 {
		uriToReq = "http://" + uriToReq
	} else {
		uriToReq = "https://" + uriToReq
	}

	if r.RequestURI != "" {
		uriToReq += r.RequestURI
	}

	transportResponse := h.transportRequest(uriToReq, w, r, target)

	if transportResponse == nil {
		go h.logHTTPRequest(httpLog.Build(target, r, nil).NoResponse())

		return
	}

	httpLog.ResponseRulesExecutionStart()

	responseIsNotSafe := response.NewResponseChecker(w, r, transportResponse, target).Handle()

	httpLog = httpLog.RequestRulesExecutionEnd()

	if responseIsNotSafe {
		go h.logHTTPRequest(httpLog.Build(target, r, nil))

		return
	}

	h.transformResponse(w, r, transportResponse)

	go h.logHTTPRequest(httpLog.Build(target, r, transportResponse))
}

//TransportRequest Transports the incoming request
func (h GuardianHandler) transportRequest(uriToReq string,
	incomingWriter http.ResponseWriter,
	incomingRequest *http.Request,
	target *models.Target) *http.Response {

	var response *http.Response
	var err error
	var req *http.Request

	//timeout is 45 secs for to pass to origin server.
	client := &http.Client{
		Timeout: time.Second * 45,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				//TODO: Check better solutions for dialcontext like timeouts.
				uri, ferr := url.Parse(addr)
				if ferr != nil {
					panic(ferr)
				}

				addr = target.OriginIPAddress + ":" + uri.Opaque

				return dialer.DialContext(ctx, network, addr)
			},
		},
	}

	req, err = http.NewRequest(incomingRequest.Method, uriToReq, incomingRequest.Body)
	for name, value := range incomingRequest.Header {
		//TODO: Do not pass the headers except whitelisted
		if name == "X-Forwarded-For" {
			continue
		}

		req.Header.Set(name, value[0])
	}

	fwIP := h.getForwardIP(incomingRequest)
	if fwIP != "" {
		req.Header.Set("X-Forwarded-For", fwIP)
	}

	response, err = client.Do(req)

	if err != nil {
		http.Error(incomingWriter, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return response
}

func (h GuardianHandler) transformResponse(incomingWriter http.ResponseWriter, incomingRequest *http.Request, response *http.Response) {
	for k, v := range response.Header {
		incomingWriter.Header().Set(k, v[0])
	}
	incomingWriter.WriteHeader(response.StatusCode)
	io.Copy(incomingWriter, response.Body)
	defer incomingRequest.Body.Close()
}

func (h GuardianHandler) logHTTPRequest(log *models.HTTPLog) {
	h.DB.LogHTTPRequest(log)
}

func (h GuardianHandler) getForwardIP(incomingRequest *http.Request) string {

	ipAddress := incomingRequest.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = incomingRequest.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = incomingRequest.RemoteAddr
	}

	return ipAddress
}
