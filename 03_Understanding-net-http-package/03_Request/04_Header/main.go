package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type mux int

func (m mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
		Header      http.Header
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var m mux
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		log.Fatalln(err)
	}
}
