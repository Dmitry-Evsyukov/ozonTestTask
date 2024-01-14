package urlGrpc

import (
	"context"
	"github.com/pkg/errors"
	"main/internal/models"
	"main/internal/url"
	urlProto "main/proto/url/gen"
)

type urlGrpcServer struct {
	urlUseCase url.UseCase
	urlProto.UnimplementedUrlServiceServer
}

func NewServer(uuc url.UseCase) urlGrpcServer {
	return urlGrpcServer{urlUseCase: uuc}
}

func (us urlGrpcServer) CreateShortAndSave(ctx context.Context, url *urlProto.Url) (*urlProto.Url, error) {
	shortUrl, err := us.urlUseCase.CreateShortUrlAndSave(ctx, models.Url{Url: url.GetUrl()})
	if err != nil {
		return nil, errors.Wrap(err, "urlGrpcServer.CreateShortAndSave.CreateShortAndSave")
	}

	return &urlProto.Url{Url: shortUrl.Url}, nil
}

func (us urlGrpcServer) GetFullUrl(ctx context.Context, url *urlProto.Url) (*urlProto.Url, error) {
	originalUrl, err := us.urlUseCase.GetFullUrl(ctx, models.Url{Url: url.GetUrl()})
	if err != nil {
		return nil, errors.Wrap(err, "urlGrpcServer.GetFullUrl.GetFullUrl")
	}

	return &urlProto.Url{Url: originalUrl.Url}, nil
}
