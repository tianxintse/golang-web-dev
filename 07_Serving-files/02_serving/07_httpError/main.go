package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", toby)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `<img src="/toby.jpg">`)
	if err != nil {
		log.Fatalln(err)
	}
}

func toby(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
		return
	}

	fmt.Println(fi.Name(), " ", fi.ModTime())
	http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)
}
