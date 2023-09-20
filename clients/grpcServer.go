package clients

import (
	"aeon-grpc/graph/model"
	"aeon-grpc/grpc"
	"aeon-grpc/interfaces"
	"context"
)

type grpcServer struct {
	grpc.UnimplementedBookServiceServer
	dbClient interfaces.StoreClient[model.Book]
}

// mustEmbedUnimplementedBookServiceServer implements grpc.BookServiceServer.
func (g grpcServer) mustEmbedUnimplementedBookServiceServer() {
	panic("unimplemented")
}

// CreateBook implements grpc.BookServiceServer.
func (g grpcServer) CreateBook(ctx context.Context, book *grpc.Book) (*grpc.Book, error) {
	newBook := model.Book{
		Isbn:    book.Isbn,
		Title:   book.Title,
		Author:  book.Author,
		Summary: book.Summary,
	}

	_, err := g.dbClient.CreateItem(newBook)
	if err != nil {
		return book, err
	}

	return book, nil
}

// DeleteBook implements grpc.BookServiceServer.
func (g grpcServer) DeleteBook(ctx context.Context, bookId *grpc.DeleteBookRequest) (*grpc.DeleteBookResponse, error) {
	key := bookId.GetId()
	var resp grpc.DeleteBookResponse
	resp.Deleted = false

	err := g.dbClient.DeleteItem(key)
	if err != nil {
		return &resp, err
	}
	resp.Deleted = true
	return &resp, nil
}

// GetBook implements grpc.BookServiceServer.
func (g grpcServer) GetBook(ctx context.Context, bookId *grpc.GetBookRequest) (*grpc.Book, error) {
	var book grpc.Book

	key := bookId.GetId()

	res, err := g.dbClient.GetItem(key)
	if err != nil {
		return &book, err
	}
	book.Title = res.Title
	book.Author = res.Author
	book.Isbn = res.Isbn
	book.Summary = res.Summary

	return &book, nil
}

// UpdateBook implements grpc.BookServiceServer.
func (g grpcServer) UpdateBook(ctx context.Context, book *grpc.Book) (*grpc.Book, error) {
	updatedBook := model.Book{
		ID:      string(book.Id),
		Title:   book.Title,
		Author:  book.Author,
		Isbn:    book.Isbn,
		Summary: book.Summary,
	}
	_, err := g.dbClient.UpdateItem(updatedBook.Isbn, updatedBook)
	if err != nil {
		return book, err
	}

	return book, nil
}

func NewGrpcServer(dbClient interfaces.StoreClient[model.Book]) grpc.BookServiceServer {
	return grpcServer{
		dbClient: dbClient,
	}
}
