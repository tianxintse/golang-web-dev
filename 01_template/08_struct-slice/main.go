package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

type car struct {
	Name  string
	Price int
}

type pc struct {
	People []person
	Cars   []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("test.html"))
}

func main() {
	p1 := person{
		First: "Tianxin",
		Last:  "Xie",
		Age:   23,
	}
	p2 := person{
		First: "Kexin",
		Last:  "Hu",
		Age:   22,
	}
	p3 := person{
		First: "Haoyue",
		Last:  "Yuan",
		Age:   22,
	}
	people := []person{p1, p2, p3}

	c1 := car{
		Name:  "Benchi",
		Price: 20000,
	}
	c2 := car{
		Name:  "Baoma",
		Price: 40000,
	}
	c3 := car{
		Name:  "QQ",
		Price: 3000,
	}
	c4 := car{
		Name:  "Mashaladi",
		Price: 120000,
	}
	cars := []car{c1, c2, c3, c4}

	data := pc{
		People: people,
		Cars:   cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
