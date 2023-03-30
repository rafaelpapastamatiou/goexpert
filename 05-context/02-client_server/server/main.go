package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
		return

	case <-ctx.Done():
		log.Println("Request cancelada pelo client")
		return
	}
}
