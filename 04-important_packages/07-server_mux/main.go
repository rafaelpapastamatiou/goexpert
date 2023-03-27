package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := make(chan int)

	fmt.Println("Starting Mux 1...")
	mux := http.NewServeMux()

	mux.HandleFunc("/", DefaultHandler)
	mux.Handle("/blog", &blog{name: "My Blog"})

	go http.ListenAndServe(":8080", mux)

	fmt.Println("Starting Mux 2...")
	mux2 := http.NewServeMux()

	mux2.HandleFunc("/", DefaultHandler2)
	mux2.Handle("/blog", &blog{name: "My Blog 2"})

	go http.ListenAndServe(":8081", mux2)

	<-c
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Mux 1"))
}

func DefaultHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Mux 2"))
}

type blog struct {
	name string
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + b.name))
}
