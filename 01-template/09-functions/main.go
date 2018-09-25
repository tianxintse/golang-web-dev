package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("test.html"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		return s
	}
	return s[:3]
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "test.html", "Hello Template!")
	if err != nil {
		log.Fatalln(err)
	}
}
