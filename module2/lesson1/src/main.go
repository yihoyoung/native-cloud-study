package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	flag.Parse()
	glog.V(4).Info("start server...")

	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	statusCode := 200
	// 1、接收客户端 request，并将 request 中带的 header 写入 response header
	for headerKey, value := range r.Header {
		w.Header().Set(headerKey, value[0])
	}
	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	// 4. 当访问 localhost/healthz 时，应返回200
	w.WriteHeader(statusCode)
	io.WriteString(w, "ok")
	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Println("request:", r.RequestURI, r.Method, statusCode)
}
