package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	p1 := doubleZero{
		person{
			Name: "James Bond",
			Age:  44,
		},
		false,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
