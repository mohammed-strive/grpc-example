package client

import (
	"aeon-grpc/utils"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateGrpcClient(caFile, serverHostOverride, serverAddress string, tls bool) *grpc.ClientConn {
	var opts []grpc.DialOption
	if tls {
		if caFile == "" {
			caFile = utils.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	fmt.Println("grpc client created")
	return conn
}
