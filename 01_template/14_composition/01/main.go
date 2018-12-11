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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  44,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
