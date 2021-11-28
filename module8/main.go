package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jw803/module2/pkg"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/healthz", pkg.WithLogging(http.HandlerFunc(healthz)))
	mux.Handle("/req2res", pkg.WithLogging(http.HandlerFunc(req2res)))

	srv := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Print("Server Started")
	<-done
	log.Print("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
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







