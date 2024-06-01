package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.URL.String())
	if request.Method != "GET" {
		byteData, _ := io.ReadAll(request.Body)
		fmt.Println(string(byteData))
	}
	fmt.Println(request.Header)

	byteData, _ := json.Marshal(Response{
		Code: 0,
		Msg:  "成功",
		Data: map[string]any{},
	})

	writer.Write(byteData)
}

func main() {

	http.HandleFunc("/index", Index)
	fmt.Println("http server running 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
