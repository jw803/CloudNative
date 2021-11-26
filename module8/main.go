package main

import (
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jw803/module2/pkg"
)

func main() {
	signalChan := make(chan os.Signal, 2) 
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		<-signalChan // 此处没有系统信号时阻塞，后续代码不执行，有信号时后续代码执行
	 
		signal.Stop(signalChan)  // 显式停止监听系统信号
		close(signalChan) // 显式关闭监听信号的通道
	}()
	
	http.Handle("/healthz", pkg.WithLogging(http.HandlerFunc(healthz)))
	http.Handle("/req2res", pkg.WithLogging(http.HandlerFunc(req2res)))
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







