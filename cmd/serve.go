package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	defaultPort     string = "8080"
	defaultPortDesc string = "Port to run the web server on"
	defaultAddress  string = "127.0.0.1"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, world!\n")
}

func swatchTimeHandler(w http.ResponseWriter, req *http.Request) {
	/* Verify Swatch Internet Time here https://beats.wiki/ */
	currentUTC := time.Now().UTC().Add(1 * time.Hour)
	beats := (currentUTC.Hour()*3600 + currentUTC.Minute()*60 + currentUTC.Second()) * 1000 / 86400
	fmt.Fprintf(w, "It is currently %d-%02d-%02d@%d", currentUTC.Year(), currentUTC.Month(), currentUTC.Day(), beats)
}

func main() {
	port := flag.String("port", defaultPort, defaultPortDesc)
	flag.Parse()
	bindAddress := fmt.Sprintf("%s:%s", defaultAddress, *port)

	mux := http.ServeMux{}
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/time", swatchTimeHandler)
	srvr := &http.Server{
		Addr:    bindAddress,
		Handler: &mux,
	}
	fmt.Printf("Starting server on http://%s", bindAddress)
	srvr.ListenAndServe()
}
