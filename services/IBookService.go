package services

import "aeon-grpc/models"

type BookService interface {
	GetBook(string) (models.Book, error)
	CreateBook(models.Book) (models.Book, error)
	UpdateBook(models.Book) (models.Book, error)
	DeleteBook(models.Book) error
}
