package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", handleHttpFunction)

	httpAddr := ":9010"
	httpsAddr := ":9011"
	certFile := "./selfsigned.crt"
	keyFile := "./selfsigned.key"

	go func() {
		log.Printf("HTTP Listening on %s...", httpAddr)
		err := http.ListenAndServe(httpAddr, nil)
		if err != nil {
			log.Printf("Error starting HTTP server: %v", err)
		}
	}()

	go func() {
		log.Printf("HTTPS Listening on %s...", httpsAddr)
		err := http.ListenAndServeTLS(httpsAddr, certFile, keyFile, nil)
		if err != nil {
			log.Printf("Error starting HTTPS server: %v", err)
		}
	}()

	select {}
}
