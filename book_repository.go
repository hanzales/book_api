package main

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type BookRepository struct {
	db *pgx.Conn
}

func NewBookRepository(db *pgx.Conn) *BookRepository {
	return &BookRepository{db}
}

func (br *BookRepository) CreateBook(book *Book) error {
	_, err := br.db.Exec(context.Background(), "INSERT INTO books (title, author) VALUES ($1, $2)", book.Title, book.Author)
	return err
}

func (br *BookRepository) GetBooks() ([]Book, error) {
	rows, err := br.db.Query(context.Background(), "SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

/*func (br *BookRepository) UpdateBook(book *Book) error {
	_, err := br.db.Exec(context.Background(), "UPDATE books SET title = $1, author = $2 WHERE id = $3", book.Title, book.Author, book.ID)
	return err
}

func (br *BookRepository) DeleteBook(bookID int) error {
	_, err := br.db.Exec(context.Background(), "DELETE FROM books WHERE id = $1", bookID)
	return err
}*/
