package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "Visitor")
}

func dog(w http.ResponseWriter, r *http.Request) {
	dogs := []string{"Labrador", "Beagle", "Bulldog"}
	tpl.ExecuteTemplate(w, "dog.gohtml", dogs)
}

func me(w http.ResponseWriter, r *http.Request) {
	name := struct {
		FirstName string
		LastName  string
	}{
		"Tianxin",
		"Xie",
	}
	tpl.ExecuteTemplate(w, "me.gohtml", name)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
