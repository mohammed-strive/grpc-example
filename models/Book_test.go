package models_test

import (
	"aeon-grpc/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

	It("should create type Book", func() {
		testTitle := "Test Title"
		book := models.NewBook("Book1", "Test Title", "Test Author", "ISBN", "This is a test book")
		Expect(book.Title).To(Equal(testTitle))
	})
})
