package main

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ocp-docs-api/internal/repo"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test"
	dbname   = "testdb"
)

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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", psqlInfo)

	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
	}

	repo := repo.New(*db)
	desc.RegisterOcpDocsApiServer(grpcServer, api.NewDocsApi(repo))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}
}
