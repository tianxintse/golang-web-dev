package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"formatDateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01_io.Copy-02-2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("test.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "test.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
