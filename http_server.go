package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/asalih/guardian/data"
	"github.com/asalih/guardian/models"

	"golang.org/x/crypto/acme/autocert"
)

/*HTTPServer The http server handler*/
type HTTPServer struct {
	AutoCertManager *autocert.Manager
}

/*NewHTTPServer HTTP server initializer*/
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{&autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("cert-cache"),
	}}
}

func (h HTTPServer) ServeHTTP() {
	srv80 := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      2 * time.Minute,
		ReadTimeout:       1 * time.Minute,
		Handler:           NewGuardianHandler(true, h.AutoCertManager),
		Addr:              ":http",
	}

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		GetCertificate:           h.certificateManager(),
		CipherSuites: []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			0xC028, //TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384
		},
	}

	srv := &http.Server{
		ReadHeaderTimeout: 40 * time.Second,
		WriteTimeout:      2 * time.Minute,
		ReadTimeout:       2 * time.Minute,
		Handler:           NewGuardianHandler(false, h.AutoCertManager),
		Addr:              ":https",
		TLSConfig:         tlsConfig,
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

		fmt.Println("Incoming TLS request:" + clientHello.ServerName)
		target := data.NewDBHelper().GetTarget(clientHello.ServerName)

		if target == nil {
			fmt.Println("Incoming TLS request: Target nil")
			return nil, err
		}

		if target.AutoCert {
			return h.AutoCertManager.GetCertificate(clientHello)
		}

		if !target.CertCrt.Valid && !target.CertKey.Valid {
			return nil, errors.New("CERTIFICATION IS NOT ENABLED")
		}

		cert, errl := h.loadCertificates(target)

		if errl != nil {
			panic(errl)
		}
		return &cert, nil
	}
}

func (h HTTPServer) loadCertificates(target *models.Target) (tls.Certificate, error) {
	return tls.X509KeyPair([]byte(target.CertCrt.String), []byte(target.CertKey.String))
}
