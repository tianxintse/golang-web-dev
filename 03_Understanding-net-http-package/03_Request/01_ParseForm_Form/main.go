package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	for k, vars := range req.Form {
		fmt.Println(k)
		for _, v := range vars {
			fmt.Println("\t", v)
		}
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var hd hotdog
	err := http.ListenAndServe(":8080", hd)
	if err != nil {
		log.Fatalln(err)
	}
}
