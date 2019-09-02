// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var addr = flag.String("addr", ":443", "http service address")

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
		if addr == "www.netsparker.com:443" {
			addr = "52.1.25.52:443"
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
		cert, _ := loadCertificates("ns.key", "ns.crt")

		return &cert, nil
	}
}

func loadCertificates(certFile, keyFile string) (tls.Certificate, error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.X509KeyPair(certPEMBlock, keyPEMBlock)
}

func loadX509KeyPair(certFile, keyFile string) (*x509.Certificate, *rsa.PrivateKey) {
	cf, e := ioutil.ReadFile(certFile)
	if e != nil {
		fmt.Println("cfload:", e.Error())
		os.Exit(1)
	}

	kf, e := ioutil.ReadFile(keyFile)
	if e != nil {
		fmt.Println("kfload:", e.Error())
		os.Exit(1)
	}
	cpb, cr := pem.Decode(cf)
	fmt.Println(string(cr))
	kpb, kr := pem.Decode(kf)
	fmt.Println(string(kr))
	crt, e := x509.ParseCertificate(cpb.Bytes)

	if e != nil {
		fmt.Println("parsex509:", e.Error())
		os.Exit(1)
	}
	key, e := x509.ParsePKCS1PrivateKey(kpb.Bytes)
	if e != nil {
		fmt.Println("parsekey:", e.Error())
		os.Exit(1)
	}
	return crt, key
}

func main() {
	flag.Parse()
	// certPEMBlock, _ := ioutil.ReadFile("Goo.pfx")
	//http.Ser(guardian)

	// if certPEMBlock != nil {
	// 	log.Fatal(certPEMBlock)
	// }

	srv := &http.Server{
		Handler: &guardianHandler{},
		Addr:    ":443",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
			GetCertificate:     getCertificate("netsparker.com"),
		},
	}

	srv.ListenAndServeTLS("", "")
	//err := http.ListenAndServeTLS(*addr, "", "", guardianHandler{})
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
