package main

import (

	// "go-extend/exnet"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	InitializeLog()
	httpservice()
}

func healthzHandle(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	ip = ip[0:strings.LastIndex(ip, ":")]
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
		}
	}
	version := os.Getenv("GOVERSION")
	w.Header().Set("content-type", "application/json")
	w.Header().Set("System-Version", version)
	w.WriteHeader(http.StatusOK)
	log.Println(ip + " " + strconv.Itoa(http.StatusOK))
}

func httpservice() {
	http.HandleFunc("/healthz", healthzHandle)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func InitializeLog() {
	log.SetPrefix("[ERROR]")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	fileName := "debug.log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
	log.SetOutput(io.MultiWriter(logFile))
}

// func getIPAddress() {
// 	ip := exnet.ClientPublicIP(r)
// 	if ip == "" {
// 		ip = exnet.ClientIP(r)
// 	}
// 	return ip
// }
