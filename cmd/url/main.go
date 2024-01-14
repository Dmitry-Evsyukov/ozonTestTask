package main

import (
	"google.golang.org/grpc"
	"log"
	"main/internal/url"
	urlGrpc "main/internal/url/delivery/grpc/server"
	urlRepository "main/internal/url/repository/in_memory"
	urlPgxRepository "main/internal/url/repository/postgres"
	urlUseCase "main/internal/url/usecase"
	"main/pkg/db/postgres"
	urlProto "main/proto/url/gen"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	EnvPostgresQueryName = "DATABASE_URL"
	EnvStorageTypeName   = "STORAGE_TYPE"
	TimeToLive           = 10 * time.Second
	TimeToLiveString     = "10 seconds"
	Port                 = 8081
)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		log.Fatalln("err while starting url micros: ", err)
	}

	var urlRepo url.Repository

	switch storageType := os.Getenv(EnvStorageTypeName); {
	case storageType == "InMemory":
		urlRepo = urlRepository.NewInMemory(TimeToLive)
	case storageType == "Postgres":
		db, err := postgres.NewPsqlDB(EnvPostgresQueryName)
		if err != nil {
			log.Fatalln("error connecting database: ", err)
		}
		defer db.Close()

		urlRepo = urlPgxRepository.NewPgRepo(db, TimeToLive, TimeToLiveString)
	default:
		log.Fatalln("error choosing data type")
	}

	urlUC := urlUseCase.New(urlRepo)
	urlManager := urlGrpc.NewServer(urlUC)
	server := grpc.NewServer()
	urlProto.RegisterUrlServiceServer(server, urlManager)

	if err = server.Serve(lis); err != nil {
		log.Fatalln("serving error: ", err)
	}
}
