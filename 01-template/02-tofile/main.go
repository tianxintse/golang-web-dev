package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("gotpl.html")

	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	err = tpl.Execute(f, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
