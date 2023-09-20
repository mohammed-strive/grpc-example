package main

import (
	"aeon-grpc/clients"
	"aeon-grpc/graph"
	"aeon-grpc/graph/model"
	pb "aeon-grpc/grpc"
	"aeon-grpc/interfaces"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const defaultGqlPort = "8080"
const defaultGrpcServerPort = "50051"

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
		port = defaultGqlPort
	}

	storeClient := clients.NewMongoClient(client)
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			Store: storeClient,
		}}),
	)

	// listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", defaultGrpcServerPort))
	// if err != nil {
	// 	log.Fatalf("unable to connect to tcp port %s: %v", defaultGrpcServerPort, err)
	// }

	// startGrpcServer(storeClient, listen)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func startGrpcServer(storeClient interfaces.StoreClient[model.Book], listen net.Listener) {
	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, clients.NewGrpcServer(storeClient))
	grpcServer.Serve(listen)
}

func startGrpcClient() pb.BookServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", defaultGrpcServerPort))
	if err != nil {
		log.Fatalf("error dailing %s: %v", defaultGrpcServerPort, err)
	}

	return pb.NewBookServiceClient(conn)
}
