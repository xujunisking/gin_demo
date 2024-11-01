package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	k := func(name string) (string, error) {
		return name + "年轻又帅气！", nil
	}

	//定义模版

	//创建一个名字是f.tmpl的模版对象，名字一定要与模版的名字能对应上
	// t := template.New("f.tmpl")
	// _, err := t.ParseFiles("./f.tmpl", "./x.templ")//这里可以解析多模版
	//t, err := template.New("f.tmpl").ParseFiles("./f.tmpl")

	t := template.New("f.tmpl")
	//告诉模板引擎，我现在多了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	//解析模版(自定义的函数必须在解析模版前设置)
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	name := "小王子"
	//渲染模版
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板(母模版写在前，子模版写在后)
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	name := "小王子"
	// 渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err%v", err)
		return
	}
}
