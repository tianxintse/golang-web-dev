package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("test.html"))
}

func main() {
	g1 := struct {
		Score1 int
		Score2 int
	}{
		Score1: 6,
		Score2: 9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
