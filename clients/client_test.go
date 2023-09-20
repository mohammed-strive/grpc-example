package clients_test

import (
	"aeon-grpc/clients"
	"aeon-grpc/graph/model"
	"aeon-grpc/interfaces"
	"context"
	"errors"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = Describe("Client", Ordered, func() {
	var client *mongo.Client
	var ctx context.Context
	var err error
	var db *mongo.Database
	var collection *mongo.Collection
	var mongoClient interfaces.StoreClient[model.Book]

	var books = []interface{}{
		model.Book{
			ID:      "ID1",
			Title:   "Test Title 1",
			Author:  "Test Author 1",
			Summary: "Test Summary 1",
			Isbn:    "abc-123",
		},
		model.Book{
			ID:      "ID2",
			Title:   "Test Title 2",
			Author:  "Test Author 2",
			Summary: "Test Summary 2",
			Isbn:    "abc-234",
		},
		model.Book{
			ID:      "ID3",
			Title:   "Test Title 3",
			Author:  "Test Author 3",
			Summary: "Test Summary 3",
			Isbn:    "abc-345",
		},
	}

	BeforeAll(func() {
		connString := "mongodb://localhost:27017"
		if connString == "" {
			log.Fatal("connecting string for MongoDB is empty")
		}
		ctx = context.TODO()
		opts := options.Client().ApplyURI(connString)
		client, err = mongo.Connect(ctx, opts)
		if err != nil {
			log.Fatalf("error connecting to local DB: %v", err)
		}
		db = client.Database("aeon")
		collection = db.Collection("books")
		_, err := collection.InsertMany(ctx, books)
		if err != nil {
			log.Fatalf("error inserting documents: %v", err)
		}

		mongoClient = clients.NewMongoClient(client)
	})

	AfterAll(func() {
		if _, err = collection.DeleteMany(ctx, bson.D{}); err != nil {
			log.Fatalf("error dropping collection 'test_books': %v", err)
		}
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("error disconnecting database: %v", err)
		}
	})

	Describe("Test GetItem() method", Ordered, func() {
		FContext("Success case", func() {
			It("should get item for valid key", func() {
				expectedResult := books[1].(model.Book)
				actualResult, err := mongoClient.GetItem(expectedResult.Isbn)
				Expect(err).To(BeNil())
				Expect(actualResult).To(Equal(expectedResult))
			})
		})
		FContext("Failure case", func() {
			It("should return error when item is not found", func() {
				_, actualErr := mongoClient.GetItem("invalid")
				expectedError := errors.New("error getting item for key invalid: mongo: no documents in result")
				Expect(actualErr).NotTo(BeNil())
				Expect(actualErr).To(Equal(expectedError))
			})
		})
	})

	Describe("Test CreateItem() method", Ordered, func() {
		FContext("Success case", func() {
			It("should create item when model is valid", func() {
				expectedResult := model.Book{
					ID:      "123456",
					Title:   "New title",
					Author:  "New Author",
					Summary: "New Summary",
					Isbn:    "abcdef123",
				}

				actualResult, err := mongoClient.CreateItem(expectedResult)
				Expect(err).To(BeNil())
				Expect(expectedResult).To(Equal(actualResult))
			})
		})
	})

	Describe("Test UpdateItem() method", Ordered, func() {
		FContext("Success case", func() {
			It("should update the item when available", func() {
				expectedResult := books[2].(model.Book)
				expectedResult.Author = "New Author 3"

				actualResult, err := mongoClient.UpdateItem(expectedResult.Isbn, expectedResult)
				Expect(err).To(BeNil())
				Expect(actualResult).To(Equal(expectedResult))
			})
		})

		FContext("Failure case", func() {
			It("should throw error when update fails", func() {
				expectedErr := errors.New("error deleting book with key test: item not found")
				actualErr := mongoClient.DeleteItem("test")
				Expect(actualErr).NotTo(BeNil())
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
	})
})
