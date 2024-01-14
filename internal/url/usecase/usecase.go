package urlUseCase

import (
	"context"
	"github.com/pkg/errors"
	"main/internal/models"
	"main/internal/url"
	httpError "main/pkg/err"
	"main/pkg/utils"
)

type urlUseCase struct {
	urlRepo url.Repository
}

func New(urlRepo url.Repository) urlUseCase {
	return urlUseCase{urlRepo: urlRepo}
}

func (u urlUseCase) CreateShortUrlAndSave(ctx context.Context, full models.Url) (models.Url, error) {
	if err := utils.ValidateStruct(ctx, full); err != nil {
		return models.Url{}, httpError.NewBadRequestError(errors.WithMessage(err, "urlUseCase.CreateShortUrlAndSave.ValidateStruct"))
	}

	shortUrl := utils.ShortenUrl(full.Url)
	if err := u.urlRepo.SaveShortToFullUrl(ctx, shortUrl, full.Url); err != nil {
		return models.Url{}, err
	}

	return models.Url{Url: shortUrl}, nil
}

func (u urlUseCase) GetFullUrl(ctx context.Context, short models.Url) (models.Url, error) {
	shortUrl, err := u.urlRepo.GetFullUrl(ctx, short.Url)
	if err != nil {
		return models.Url{}, err
	}

	return models.Url{Url: shortUrl}, nil
}
