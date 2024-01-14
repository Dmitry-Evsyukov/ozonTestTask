package urlPgxRepository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"main/internal/url"
	"time"
)

type urlRepo struct {
	db           *sqlx.DB
	ttlString    string
	timeToLive   time.Duration
	stopCleaning chan struct{}
}

func NewPgRepo(db *sqlx.DB, ttlDur time.Duration, ttlString string) url.Repository {
	result := urlRepo{db: db,
		ttlString:    ttlString,
		timeToLive:   ttlDur,
		stopCleaning: make(chan struct{})}
	go result.cleanUp()
	return result
}

func (im urlRepo) SaveShortToFullUrl(ctx context.Context, short, full string) error {
	_, err := im.db.ExecContext(ctx, SaveShortAndFullQuery, full, short)
	if err != nil {
		return errors.WithMessage(err, "urlRepo.SaveShortToFullUrl.ExecContext")
	}

	return nil
}

func (im urlRepo) GetFullUrl(ctx context.Context, short string) (string, error) {
	originalUrl := ""
	err := im.db.QueryRowContext(ctx, GetFullUrlByShortQuery, short).Scan(&originalUrl)
	if err != nil {
		return "", errors.Wrap(err, "urlRepo.GetFullUrl.QueryRowContext")
	}

	return originalUrl, nil
}

func (im urlRepo) cleanUp() {
	ticker := time.NewTicker(im.timeToLive)
	for {
		select {
		case <-ticker.C:
			_, err := im.db.Exec(CleanUpOldRecords)
			if err != nil {
				_ = fmt.Errorf("error cleaning postgres: %v", err)
			}
		case <-im.stopCleaning:
			ticker.Stop()
			return
		}
	}
}

func (im urlRepo) stopCleanUp() {
	im.stopCleaning <- struct{}{}
}
