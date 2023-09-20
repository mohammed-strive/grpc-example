package clients

import (
	"aeon-grpc/graph/model"
	"aeon-grpc/interfaces"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type client struct {
	mc *mongo.Client
}

func NewMongoClient(mc *mongo.Client) interfaces.StoreClient[model.Book] {
	return client{
		mc: mc,
	}
}

func (c client) GetItem(key string) (model.Book, error) {
	var book model.Book

	collection := c.mc.Database("aeon").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := bson.D{{Key: "isbn", Value: string(key)}}
	err := collection.FindOne(ctx, query).Decode(&book)
	if err != nil {
		return book, fmt.Errorf("error getting item for key %s: %v", key, err)
	}

	return book, nil
}

func (c client) CreateItem(newBook model.Book) (model.Book, error) {

	collection := c.mc.Database("aeon").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, newBook)
	if err != nil {
		return newBook, fmt.Errorf("error inserting document: %v", err)
	}

	return newBook, nil
}

func (c client) UpdateItem(key string, book model.Book) (model.Book, error) {

	collection := c.mc.Database("aeon").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "isbn", Value: key}}

	_, err := collection.ReplaceOne(ctx, filter, book)
	if err != nil {
		return book, fmt.Errorf("error updating: %v", err)
	}

	return book, nil
}

func (c client) DeleteItem(key string) error {

	collection := c.mc.Database("aeon").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "isbn", Value: key}}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting book with key %s: %v", key, err)
	}
	if res.DeletedCount != 1 {
		return fmt.Errorf("error deleting book with key %s: item not found", key)
	}

	return nil
}
