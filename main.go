package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 定义处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// 启动服务器
	port := ":8090"
	log.Printf("服务器启动在端口 %s", port)
	log.Printf("访问 http://localhost%s 查看结果", port)

	// 启动HTTP服务器
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
