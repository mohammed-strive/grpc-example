package repository

import "aeon-grpc/models"

type BooksRepository interface {
	GetBook(string) (models.Book, error)
	AddBook(models.Book) (models.Book, error)
	UpdateBook(models.Book) (models.Book, error)
	RemoveBook(models.Book) error
}
