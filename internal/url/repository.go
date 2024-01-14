package url

import "context"

type Repository interface {
	SaveShortToFullUrl(ctx context.Context, short, full string) error
	GetFullUrl(ctx context.Context, short string) (string, error)
}
