package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
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
