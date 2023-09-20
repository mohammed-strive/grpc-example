package models

type Book struct {
	Id, Title, Author, ISBN, Summary string
}

func NewBook(id, title, author, isbn, summary string) Book {
	return Book{
		Id:      id,
		Title:   title,
		Author:  author,
		ISBN:    isbn,
		Summary: summary,
	}
}
