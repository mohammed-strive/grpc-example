package server

import (
	"fmt"
	"log"
	"net"

	"aeon-grpc/graph/model"
	pb "aeon-grpc/grpc"
	"aeon-grpc/grpc/services"
	"aeon-grpc/interfaces"
	"aeon-grpc/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func StartGrpcServer(tls bool, address, certFile, keyFile string, dbClient interfaces.StoreClient[model.Book]) {
	var options []grpc.ServerOption

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to start grpc server at %s: %v", address, err)
	}
	if tls {
		if certFile == "" {
			certFile = utils.Path("x509/server_cert.pem")
		}
		if keyFile == "" {
			keyFile = utils.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		options = []grpc.ServerOption{grpc.Creds(creds)}
	}

	server := grpc.NewServer(options...)
	service := services.NewGrpcBookService(dbClient)
	pb.RegisterBookServiceServer(server, service)
	fmt.Println("Started grpc server")
	server.Serve(listen)
}
