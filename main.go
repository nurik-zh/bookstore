package main

import (
	"net/http"

	"bookstore/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Роуты для книг
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Роуты для авторов
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors", handlers.AddAuthor).Methods("POST")

	// Роуты для категорий
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", handlers.AddCategory).Methods("POST")

	http.ListenAndServe(":8080", r)
}
