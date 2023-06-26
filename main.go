package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	caPEM, err := ioutil.ReadFile("tls/ca.pem")
	if err != nil {
		log.Fatal("err reading ca.pem:", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPEM) {
		log.Fatal("err adding CA to cert pool")
	}

	config := tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}

	srv := http.Server{
		Addr:      ":8081",
		TLSConfig: &config,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.TLS.PeerCertificates) > 0 {
			cns := make([]string, len(r.TLS.PeerCertificates))
			for i, cert := range r.TLS.PeerCertificates {
				cns[i] = cert.Subject.CommonName
			}

			w.Write([]byte(fmt.Sprintf("Hello, %v!", strings.Join(cns, ", "))))
		} else {
			w.Write([]byte("Hello, world!"))
		}
	})

	log.Println("listen and server on", srv.Addr)
	if err := srv.ListenAndServeTLS("tls/server.pem", "tls/server-key.pem"); err != nil {
		log.Fatal("err ListenAndServe:", err)
	}
}
