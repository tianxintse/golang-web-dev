package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "gotpl.html",
		`Helping others is helping yourself.`)
	if err != nil {
		log.Fatalln(err)
	}
}
