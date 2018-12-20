package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", "Visitor")
}
