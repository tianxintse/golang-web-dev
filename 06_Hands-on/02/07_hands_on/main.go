package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	var method, uri string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())
		method = fs[0]
		uri = fs[1]
		fmt.Printf("METHOD: %s, URI: %s", method, uri)
		break
	}
	body := "METHOD: " + method + ",URI: " + uri
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "%s\r\n", body)
}
