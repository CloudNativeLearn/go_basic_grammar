package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 单独写回调函数
	http.HandleFunc("/go",myHandler)
	// addr:监听地址
	// handler:回调函数
	http.ListenAndServe("127.0.0.1:8000",nil)
}

// handler函数
func myHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println(r.RemoteAddr,"链接成功")
	// 请求方式 Get  Post  Delete Put  Update
	fmt.Println("method:",r.Method)
	// go
	fmt.Println("url:",r.URL.Path)
	fmt.Println("header:",r.Header)
	fmt.Println("body",r.Body)
	//回复
	w.Write([]byte("www.51mh.com"))
}
