# grpc-example

go + GraphQL + grpc example

This project demonstrates the use of GraphQl and grpc for creating services in
golang.

- The grpc server runs at localhost:50051
- The graphql playground runs at localhost:8080

### Requirements

- [mongodb](https://www.mongodb.com) (either running in cloud or locally. You
  have to give the connection string as ENV).
- [golang](https://go.dev) (anything above 1.18 should work).
- [protoc](https://protobuf.dev) (protocol buffer compiler.)

### Running the project

- Pull the repo.
- Set the `MONGODB_URI` ENV var to a running instance of MongoDB.
- Run `go mod tidy`
- Run go run main.go
- Navigate to `http://localhost:8080` to use the graphql playground.
