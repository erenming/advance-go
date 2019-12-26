package main

import (
	"io"
	"log"
	"net/http"
)

func sayHello(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(200)
	io.WriteString(wr, "hello, world")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}