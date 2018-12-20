package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `<img src="/resources/toby.jpg">`)
	if err != nil {
		log.Fatalln(err)
	}
}
