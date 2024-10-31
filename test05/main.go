package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 结构体字段显示，必须首字母大写
type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	//渲染模版
	u1 := User{
		Name:   "小丸子",
		Gender: "男",
		Age:    18,
	}

	m1 := map[string]interface{}{
		"name":   "小猫猫",
		"gender": "女",
		"age":    16,
	}

	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}
