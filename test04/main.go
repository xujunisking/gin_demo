package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//遇事不决 先写注释
//1.定义模版

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
		return
	}

	//3.渲染模版
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}
