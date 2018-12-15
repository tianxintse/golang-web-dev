package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcom to the index page.")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a page for your dogs.")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My name is Tianxin Xie. It's nice to have you here.")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
