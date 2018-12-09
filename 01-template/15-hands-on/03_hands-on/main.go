package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		{
			Name:    "name 1",
			Address: "address 1",
			City:    "city 1",
			Zip:     "10001",
			Region:  "Southern",
		},
		{
			Name:    "name 2",
			Address: "address 2",
			City:    "city 2",
			Zip:     "10002",
			Region:  "Central",
		},
		{
			Name:    "name 3",
			Address: "address 3",
			City:    "city 3",
			Zip:     "10003",
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
