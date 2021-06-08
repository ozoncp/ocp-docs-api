package main

import (
	"flag"
	"fmt"
	"github.com/ocp-docs-api/internal/api"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main () {
	fmt.Println("Start gRPC server")

	const grpcPort = ":7002"

	var grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpDocsApiServer(s, api.NewDocsApi())

	fmt.Printf("Server listening on %s\n", *grpcEndpoint)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
