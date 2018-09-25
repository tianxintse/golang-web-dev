package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("test.html"))
}

func main() {
	a := map[string]string{
		"America": "English",
		"China":   "Chinese",
		"Japan":   "Japanese",
		"France":  "French",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "test.html", a)
	if err != nil {
		log.Fatalln(err)
	}
}
