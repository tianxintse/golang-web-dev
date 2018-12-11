package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatalln(err)
	}
}
