package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func main() {
	connString := "postgresql://postgres:postgres@localhost:5432/bookstore"
	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	bookRepo := NewBookRepository(db)
	bookService := NewBookService(bookRepo)
	bookController := NewBookController(bookService)

	r := mux.NewRouter()

	// Kitapları eklemek için POST isteği
	r.HandleFunc("/books", bookController.CreateBook).Methods("POST")

	// Kitapları listelemek için GET isteği
	r.HandleFunc("/books", bookController.GetBooks).Methods("GET")

	/*	// Kitapları güncellemek için PUT isteği
		r.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")

		// Kitapları silmek için DELETE isteği
		r.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")*/

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
