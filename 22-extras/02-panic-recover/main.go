package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func panicRecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic recover middleware started")

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				debug.PrintStack()
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("This is a panic!")
	})

	println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", panicRecoverMiddleware(mux)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
