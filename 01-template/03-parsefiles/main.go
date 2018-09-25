package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("gotpl1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("gotpl2.txt")
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("gotpl3.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute gotpl1.txt
	err = tpl.ExecuteTemplate(os.Stdout, "gotpl1.txt", nil)
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
