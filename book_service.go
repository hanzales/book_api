package main

type BookService struct {
	repository *BookRepository
}

func NewBookService(repository *BookRepository) *BookService {
	return &BookService{repository}
}

func (bs *BookService) CreateBook(book *Book) error {
	return bs.repository.CreateBook(book)
}

func (bs *BookService) GetBooks() ([]Book, error) {
	return bs.repository.GetBooks()
}

/*func (bs *BookService) UpdateBook(book *Book) error {
return bs.repository.UpdateBook(book)
}

func (bs *BookService) DeleteBook(bookID int) error {
return bs.repository.DeleteBook(bookID)
}*/
