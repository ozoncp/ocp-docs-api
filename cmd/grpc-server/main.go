package main

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ocp-docs-api/internal/api"
	"github.com/ocp-docs-api/internal/metrics"
	"github.com/ocp-docs-api/internal/producer"
	"github.com/ocp-docs-api/internal/repo"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
)

const (
	grpcPort  = 7002
	chunkSize = 5
)

const (
	dbHost     = "database"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "postgres"
)

const (
	kafkaBroker = "kafka:9092"
)

const (
	metricsPort   = ":9101"
    metricsHandle = "/metrics"
)

func main() {
	var err error
	go func() {
		if err = runMetrics(); err != nil {
			log.Fatal().Msgf("failed to run metrics: %v", err)
			return
		}
	}()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err = runGrpc()
	if err != nil {
		log.Fatal().Err(err)
	}
}

func runGrpc() error {
	grpcEndpoint := fmt.Sprintf(":%d", grpcPort)
	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start feedback grpc server at %v", grpcEndpoint)
		return err
	}
	log.Info().Msgf("Starting server at %v...", grpcEndpoint)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	db, err := connectToDataBase()
	if err != nil {
		log.Fatal().Err(err).Msgf("Issue with database connection")
		return err
	}

	brokerProducer, err := connectToBroker()

	if err != nil {
		log.Fatal().Err(err).Msgf("Issue with broker connection")
		return err
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal().Msgf("failed to close db connection: %v", err)
		}
		if err = brokerProducer.Close(); err != nil {
			log.Fatal().Msgf("failed to close producer connection: %v", err)
		}
	}()

	repo := repo.New(*db, chunkSize)
	desc.RegisterOcpDocsApiServer(grpcServer, api.NewDocsApi(repo, brokerProducer))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}

	return nil
}

func runMetrics() error {
	metrics.RegisterMetrics()
	http.Handle(metricsHandle, promhttp.Handler())
	return http.ListenAndServe(metricsPort, nil)
}

func connectToDataBase() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sqlx.Open("pgx", psqlInfo)
	if err != nil {
		log.Error().Err(err).Msgf("db can't be open")
		return db, nil
	}
	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
		return db, err
	}
	return db, nil
}

func connectToBroker() (producer.Producer, error) {
	kafkaBroker := []string{kafkaBroker}
	prod, err := producer.NewProducer(kafkaBroker,"OcpDocsApi")
	if err != nil {
		log.Error().Err(err).Msgf("failed to create kafka producer")
		return nil, err
	}
	return prod, nil
}
