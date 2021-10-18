package main

import (
	"net/http"
	"os"
	"strings"
	"fmt"
	"net"
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


type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func WithLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         200,
		}

		ip, err := getIP(r)

		if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("No valid ip"))
		}

		h.ServeHTTP(recorder, r)
		fmt.Printf("ClientId: %s\n", ip)
		fmt.Printf("Http Status: %d\n", recorder.Status)
	})
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
			return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
			netIP := net.ParseIP(ip)
			if netIP != nil {
					return ip, nil
			}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
			return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
			return ip, nil
	}
	
	return "", fmt.Errorf("No valid ip found")
}

