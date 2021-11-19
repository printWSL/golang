package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v \n", err)
	}

	msg := "这是index页面"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v \n", err)
	}

	msg := "这是home页面"
	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start faield, err %v", err)
	}
}
