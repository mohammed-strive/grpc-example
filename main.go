package main

import (
	"aeon-grpc/clients"
	"aeon-grpc/graph"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Could not find .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalf("Mongodb URI not set. Set the URI in MONGODB_URI env val")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("unable to establish connection to DB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("unable to disconnect from DB: %v", err)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	storeClient := clients.NewMongoClient(client)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Store: storeClient,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
