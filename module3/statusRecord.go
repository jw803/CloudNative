package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

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