package services

import (
	"aeon-grpc/models"
	"aeon-grpc/repository"
	"log"
)

// Run command to generate / update mock
// mockgen -source=./IBookService.go -destination=../mocks/BookService.mock.go

type bookService struct {
	logger *log.Logger
	repo   repository.BooksRepository
}

func NewBookService(logger *log.Logger, bookRepo repository.BooksRepository) BookService {
	return bookService{
		logger: logger,
		repo:   bookRepo,
	}
}

func (b bookService) GetBook(bookId string) (models.Book, error) {
	panic("not implemented")
}

func (b bookService) CreateBook(newBook models.Book) (models.Book, error) {
	panic("not implemented")
}

func (b bookService) UpdateBook(book models.Book) (models.Book, error) {
	panic("not implemented")
}

func (b bookService) DeleteBook(book models.Book) error {
	panic("not implemented")
}
