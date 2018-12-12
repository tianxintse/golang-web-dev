package main

import (
	"io"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		_, err := io.WriteString(w, "doggy doggy doggy")
		if err != nil {
			log.Fatalln(err)
		}
	case "/cat":
		_, err := io.WriteString(w, "kitty kitty kitty")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatalln(err)
	}
}
