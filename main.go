// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var addr = flag.String("addr", ":443", "http service address")
var db = &DBHelper{}

var dialer = &net.Dialer{
	Timeout:   30 * time.Second,
	KeepAlive: 30 * time.Second,
	DualStack: true,
}

type guardianHandler struct{}

func (h guardianHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isNotSafeRequest := lookRequest(r)

	if isNotSafeRequest {
		fmt.Fprintf(w, "Not nice. %s", r.URL.Path)

		return
	}

	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {

		uri, ferr := url.Parse(addr)
		if ferr != nil {
			panic(ferr)
		}

		target := db.GetTarget(uri.Scheme)

		if target != nil {
			addr = target.OriginIPAddress + ":" + strconv.Itoa(target.Port)
		}

		return dialer.DialContext(ctx, network, addr)
	}

	resp, _ := http.Get("https://" + r.Host)
	defer resp.Body.Close()

	for name, values := range resp.Header {
		w.Header()[name] = values
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)

	//fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func lookRequest(r *http.Request) bool {
	return strings.Contains(r.URL.RawQuery, "%3Cscript%3E")
}

func getCertificate(arg interface{}) func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	var err error
	// if host, ok := arg.(string); ok {

	// } else if o, ok := arg.(Certopts); ok {
	// 	opts = o
	// } else {
	// 	err = errors.New("Invalid arg type, must be string(hostname) or Certopt{...}")
	// }
	return func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if err != nil {
			return nil, err
		}

		target := db.GetTarget(clientHello.ServerName)

		if target == nil {
			return nil, err
		}

		cert, err := loadCertificates(target)

		if err != nil {
			panic(err)
		}
		return &cert, nil
	}
}

// func loadCertificates(certFile, keyFile string) (tls.Certificate, error) {
// 	certPEMBlock, err := ioutil.ReadFile(certFile)
// 	fmt.Println(string(certPEMBlock))
// 	if err != nil {
// 		return tls.Certificate{}, err
// 	}
// 	keyPEMBlock, err := ioutil.ReadFile(keyFile)
// 	fmt.Println(string(keyPEMBlock))
// 	if err != nil {
// 		return tls.Certificate{}, err
// 	}
// 	return tls.X509KeyPair(certPEMBlock, keyPEMBlock)
// }

func loadCertificates(target *Target) (tls.Certificate, error) {
	return tls.X509KeyPair([]byte(target.CertCrt), []byte(target.CertKey))
}

func main() {
	flag.Parse()

	srv := &http.Server{
		Handler: &guardianHandler{},
		Addr:    ":443",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: false,
			GetCertificate:     getCertificate("netsparker.com"),
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
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
		},
	}

	srv.ListenAndServeTLS("", "")
	//err := http.ListenAndServeTLS(*addr, "", "", guardianHandler{})
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
