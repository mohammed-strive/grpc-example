package repository

import (
	"aeon-grpc/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type booksRepository struct {
	logger *log.Logger
	client *mongo.Client
}

// AddBook implements BooksRepository.
func (booksRepository) AddBook(models.Book) (models.Book, error) {
	panic("unimplemented")
}

// GetBook implements BooksRepository.
func (booksRepository) GetBook(string) (models.Book, error) {
	panic("unimplemented")
}

// RemoveBook implements BooksRepository.
func (booksRepository) RemoveBook(models.Book) error {
	panic("unimplemented")
}

// UpdateBook implements BooksRepository.
func (booksRepository) UpdateBook(models.Book) (models.Book, error) {
	panic("unimplemented")
}

func NewBooksRepository(logger *log.Logger, client *mongo.Client) BooksRepository {
	return booksRepository{
		logger: logger,
		client: client,
	}
}
