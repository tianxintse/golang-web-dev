package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.txt"))
}

func main() {
	// Execute gotpl1.txt
	err := tpl.ExecuteTemplate(os.Stdout, "gotpl1.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute gotpl2.txt
	err = tpl.ExecuteTemplate(os.Stdout, "gotpl2.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute gotpl3.txt
	err = tpl.ExecuteTemplate(os.Stdout, "gotpl3.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
