package main

import (
	"aeon-grpc/clients"
	"aeon-grpc/constants"
	"aeon-grpc/graph"
	pb "aeon-grpc/grpc"
	gclient "aeon-grpc/grpc/client"
	"aeon-grpc/grpc/server"
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile           = flag.String("cert_file", "", "The TLS cert file")
	keyFile            = flag.String("key_file", "", "The TLS key file")
	address            = flag.String("port", "localhost:50051", "The server port")
	serverHostOverride = flag.String(
		"server_host_override",
		"x.test.example.com",
		"The server name used to verify the hostname returned by the TLS handshake",
	)
)

func main() {
	flag.Parse()

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
		port = constants.DefaultMongoPort
	}

	grpcClient := gclient.CreateGrpcClient(*certFile, *serverHostOverride, *address, *tls)
	defer grpcClient.Close()
	bookServiceClient := pb.NewBookServiceClient(grpcClient)

	storeClient := clients.NewMongoClient(client)
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			GqlClient: bookServiceClient,
		}}),
	)

	go server.StartGrpcServer(*tls, *address, *certFile, *keyFile, storeClient)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
