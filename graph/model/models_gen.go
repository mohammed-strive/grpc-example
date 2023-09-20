// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Isbn    string `json:"isbn"`
	Summary string `json:"summary"`
}

type BookInput struct {
	Title   *string `json:"title,omitempty"`
	Author  *string `json:"author,omitempty"`
	Isbn    *string `json:"isbn,omitempty"`
	Summary *string `json:"summary,omitempty"`
}