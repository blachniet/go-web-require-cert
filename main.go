package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
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
		w.Write([]byte("Hello World"))
	})

	log.Println("listen and server on", srv.Addr)
	if err := srv.ListenAndServeTLS("tls/server.pem", "tls/server-key.pem"); err != nil {
		log.Fatal("err ListenAndServe:", err)
	}
}
