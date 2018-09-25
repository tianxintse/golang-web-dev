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
	a := []string{"Apple Inc", "Huawei Inc", "Alibaba Group", "Microsoft"}

	err := tpl.ExecuteTemplate(os.Stdout, "test.html", a)
	if err != nil {
		log.Fatalln(err)
	}
}
