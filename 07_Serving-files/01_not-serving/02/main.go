package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `
	<!--not serving from our server-->
	<img src="/toby.jpg">
	`) // this will not serve, because we have no handler for it
	if err != nil {
		log.Fatalln(err)
	}
}
