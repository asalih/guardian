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
	DB          *data.DBHelper
	CertManager *autocert.Manager
}

//var CertManagerHTTPHandler =

/*NewHTTPServer HTTP server initializer*/
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{&data.DBHelper{}, &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("cert-cache"),
		HostPolicy: autocert.HostWhitelist("guardsparker.com", "www.guardsparker.com"),
	}}
}

func (h HTTPServer) ServeHTTP() {

	srv80 := &http.Server{
		//ReadHeaderTimeout: 20 * time.Second,
		//WriteTimeout:      2 * time.Minute,
		//ReadTimeout:       1 * time.Minute,
		Handler: NewGuardianHandler(true, h.CertManager),
		Addr:    ":http",
	}

	tlsConfig := &tls.Config{
		GetCertificate: h.CertManager.GetCertificate,
	}

	srv := &http.Server{
		//ReadHeaderTimeout: 40 * time.Second,
		//WriteTimeout:      2 * time.Minute,
		//ReadTimeout:       2 * time.Minute,
		Handler:   NewGuardianHandler(false, h.CertManager),
		Addr:      ":https",
		TLSConfig: tlsConfig,
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
		target := h.DB.GetTarget(clientHello.ServerName)

		if target == nil {
			fmt.Println("Incoming TLS request: Target nil")
			return nil, err
		}

		if target.AutoCert {
			fmt.Println("AutoCert GetCertificate triggered.")
			leCert, lerr := h.CertManager.GetCertificate(clientHello)

			fmt.Println(leCert)
			fmt.Println(lerr)

			return leCert, lerr
		}

		if !target.CertCrt.Valid && !target.CertKey.Valid {
			return nil, errors.New("Certification is not enabled.")
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
