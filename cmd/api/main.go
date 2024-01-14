package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"main/internal/middleware"
	grpcUrl "main/internal/url/delivery/grpc/client"
	urlDelivery "main/internal/url/delivery/http"
	urlProto "main/proto/url/gen"
)

const Address = ":8082"

func main() {
	logger := logrus.New()
	urlServiceConn, err := grpc.Dial("url:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error connecting to url micros ", err)
	}

	urlAgent := grpcUrl.NewClient(urlProto.NewUrlServiceClient(urlServiceConn))
	urlHandler := urlDelivery.NewHandler(urlAgent)

	s := echo.New()
	middlewareManager := middleware.NewManager(logger)

	urlDelivery.MapUrlRoutes(s, urlHandler, middlewareManager)
	if err := s.Start(Address); err != nil {
		log.Fatalln("error starting api server: ", err)
	}
}
