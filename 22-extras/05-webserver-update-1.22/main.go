package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", ListBooksHandler)
	mux.HandleFunc("GET /books/{id}", GetBookHandler)
	mux.HandleFunc("POST /books", CreateBookHandler)
	mux.HandleFunc("PUT /books/{id}", UpdateBookHandler)
	mux.HandleFunc("DELETE /books/{id}", DeleteBookHandler)

	mux.HandleFunc("GET /books/dir/{dir...}", GetBookDirHandler)

	mux.HandleFunc("GET /precedence/{id}", GetPrecedenceGeneralHandler)
	mux.HandleFunc("GET /precedence/xxx", GetPrecedenceSpecificHandler)

	http.ListenAndServe(":8080", mux)
}

func ListBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Books"))
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Getting Book! ID: " + id))
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating Book!"))
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Updating Book! ID: " + id))
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Deleting Book! ID: " + id))
}

func GetBookDirHandler(w http.ResponseWriter, r *http.Request) {
	dir := r.PathValue("dir")
	w.Write([]byte("Getting Book Directory! Dir: " + dir))
}

func GetPrecedenceSpecificHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Specific Precedence Handler"))
}

func GetPrecedenceGeneralHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("General Precedence Handler"))
}
