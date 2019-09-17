package main

import (
	"crypto/tls"
	"net/http"
)

/*HTTPServer The http server handler*/
type HTTPServer struct {
	DB *DBHelper
}

/*NewHTTPServer HTTP server initializer*/
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{&DBHelper{}}
}

func (h HTTPServer) ServeHTTP() {
	srv80 := &http.Server{
		Handler: NewGuardianHandler(true),
		Addr:    ":80",
	}

	srv := &http.Server{
		Handler: NewGuardianHandler(false),
		Addr:    ":443",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: false,
			GetCertificate:     h.certificateManager(),
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_RC4_128_SHA,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
			},
			PreferServerCipherSuites: true,
		},
	}

	go srv80.ListenAndServe()
	srv.ListenAndServeTLS("", "")
}

func (h HTTPServer) certificateManager() func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	var err error

	return func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if err != nil {
			return nil, err
		}

		target := h.DB.GetTarget(clientHello.ServerName)

		if target == nil {
			return nil, err
		}

		cert, errl := h.loadCertificates(target)

		if errl != nil {
			panic(errl)
		}
		return &cert, nil
	}
}

func (h HTTPServer) loadCertificates(target *Target) (tls.Certificate, error) {
	return tls.X509KeyPair([]byte(target.CertCrt), []byte(target.CertKey))
}
