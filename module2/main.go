package main

import (
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/healthz", WithLogging(http.HandlerFunc(healthz)))
	http.Handle("/req2res", WithLogging(http.HandlerFunc(req2res)))
	http.ListenAndServe(":3000", nil)
}

func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func req2res(w http.ResponseWriter, req *http.Request) {
	reqHeader := req.Header.Clone()

	for k, v := range reqHeader {
		w.Header().Add(k, strings.Join(v, ""))
	}
	version := os.Getenv("VERSION")
	w.Header().Add("version", version)
	w.WriteHeader(http.StatusOK)
}

