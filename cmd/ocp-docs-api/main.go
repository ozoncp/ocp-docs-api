package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/ocp-docs-api/internal/api"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var grpcPort int

func init() {
	flag.IntVar(&grpcPort, "port", 1235, "GRPC server port")
}

func main() {
	flag.Parse()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	grpcEndpoint := fmt.Sprintf("localhost:%d", grpcPort)

	lis, err := net.Listen("tcp", grpcEndpoint)

	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start feedback grpc server at %v", grpcEndpoint)
	}

	log.Info().Msgf("Starting server at %v...", grpcEndpoint)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	desc.RegisterOcpDocsApiServer(grpcServer, api.NewDocsApi())

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}
}