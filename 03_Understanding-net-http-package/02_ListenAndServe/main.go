package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Go HTPP Server!\n")
}

func main() {
	var hd hotdog
	http.ListenAndServe(":8080", hd)
}
