package grpcUrl

import (
	"context"
	"github.com/pkg/errors"
	"main/internal/models"
	urlProto "main/proto/url/gen"
)

type client struct {
	urlManager urlProto.UrlServiceClient
}

func NewClient(um urlProto.UrlServiceClient) client {
	return client{urlManager: um}
}

func (c client) CreateShortUrlAndSave(ctx context.Context, url models.Url) (models.Url, error) {
	shortUrl, err := c.urlManager.CreateShortAndSave(ctx, &urlProto.Url{
		Url: url.Url,
	})
	if err != nil {
		return models.Url{}, errors.Wrap(err, "urlGrpcClient.CreateShortUrlAndSave.CreateShortAndSave")
	}

	return models.Url{Url: shortUrl.GetUrl()}, nil
}

func (c client) GetFullUrl(ctx context.Context, url models.Url) (models.Url, error) {
	originalUrl, err := c.urlManager.GetFullUrl(ctx, &urlProto.Url{
		Url: url.Url,
	})
	if err != nil {
		return models.Url{}, errors.Wrap(err, "urlGrpcClient.CreateShortUrlAndSave.GetFullUrl")
	}

	return models.Url{Url: originalUrl.GetUrl()}, nil
}
