package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 使用 GET 方法注册处理器
	mux := http.NewServeMux()
	mux.HandleFunc("GET /goods/{name...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "iPhone")
		log.Println(r.PathValue("name"))
	})
	// 创建服务器
	server := http.Server{
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
