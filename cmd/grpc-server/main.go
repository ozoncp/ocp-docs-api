package main

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ocp-docs-api/internal/api"
	"github.com/ocp-docs-api/internal/producer"
	"github.com/ocp-docs-api/internal/repo"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

const (
	grpcPort = 82
	chunkSize = 5
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "test"
	dbname   = "postgres"
)

func runGrpc() error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	grpcEndpoint := fmt.Sprintf("0.0.0.0:%d", grpcPort)

	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start feedback grpc server at %v", grpcEndpoint)
		return err
	}
	log.Info().Msgf("Starting server at %v...", grpcEndpoint)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", psqlInfo)
	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
		return nil
	}

	repo := repo.New(*db, chunkSize)
	prod, err := producer.NewProducer("OcpDocsApi")
	if err != nil {
		log.Error().Err(err).Msgf("failed to create kafka producer")
	}
	desc.RegisterOcpDocsApiServer(grpcServer, api.NewDocsApi(repo, prod))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}

	return nil
}

func main() {
	err := runGrpc()
	if err != nil {
		log.Fatal().Err(err)
	}
}
