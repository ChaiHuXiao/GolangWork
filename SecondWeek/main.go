package main

import (
	"fmt"
	"go-extend/exnet"
	"net/http"
	"os"
	"path"
)

func main() {
	toLog()
	// fmt.Println(prepareLog())
}

func toLog() {
	getwd, _ := os.Getwd() // 设置文件存储路径在当前工作目录下
	newPath := path.Join(getwd, "./exercise/logger/logFile")
	InitializeLogger(newPath)
	// for {
	// 	my.Debug("这是一条Debug信息")
	// 	time.Sleep(time.Second)
	// 	my.TRACE("这是一条Trace信息")
	// 	time.Sleep(time.Second)
	// 	my.INFO("这是一条Info信息")
	// 	time.Sleep(time.Second)
	// 	my.WARNING("这是一条Warning信息")
	// 	time.Sleep(time.Second)
	// 	my.ERROR("这是一条Error信息")
	// 	time.Sleep(time.Second)
	// 	my.FATAL("这是一条Fatal信息")
	// 	time.Sleep(time.Second)
	// }
}

func healthzHandle(w http.ResponseWriter, r *http.Request) {
	// print(r.RemoteAddr)
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	println(ip)
	// log.Logger
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
			// println(len(v))
			// // println("HeaderKey: " + k) // + " HeaderValue: " + v[0] + " HeaderValueLength: " + strconv.Itoa(len(v)))
		}
	}
	version := os.Getenv("GOVERSION")
	fmt.Printf("系统版本：%s", version)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("System-Version", version)
	w.WriteHeader(http.StatusOK)

}

func httpservice() {
	http.HandleFunc("/healthz", healthzHandle)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
