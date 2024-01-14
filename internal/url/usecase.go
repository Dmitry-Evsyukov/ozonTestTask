package url

import (
	"context"
	"main/internal/models"
)

type UseCase interface {
	CreateShortUrlAndSave(ctx context.Context, url models.Url) (models.Url, error)
	GetFullUrl(ctx context.Context, url models.Url) (models.Url, error)
}
