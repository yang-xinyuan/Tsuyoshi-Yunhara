package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 处理请求头
		for key, values := range r.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// 处理版本号
		version := os.Getenv("VERSION")
		if version != "" {
			w.Header().Set("VERSION", version)
		}

		// 记录访问日志
		log.Printf("Client IP: %s, Status Code: %d\n", r.RemoteAddr, http.StatusOK)

		// 处理 /healthz 路径
		if r.URL.Path == "/healthz" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Server is healthy")
			return
		}

		// 处理其他路径
		http.NotFound(w, r)
	})

	log.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

