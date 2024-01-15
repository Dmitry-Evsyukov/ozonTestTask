package urlPgxRepository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUrlsRepo_SaveShortToFull(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()
	const ttl = 1 * time.Second
	const ttlStr = "1 seconds"

	urlRepo := NewPgRepo(sqlxDB, ttl, ttlStr)

	t.Run("Success", func(t *testing.T) {
		originalUrl := "http://example.com"
		shortUrl := "mockRecord"

		mock.ExpectExec(SaveShortAndFullQuery).WithArgs(originalUrl, shortUrl).WillReturnResult(sqlmock.NewResult(1, 1))

		err = urlRepo.SaveShortToFullUrl(context.Background(), shortUrl, originalUrl)

		require.NoError(t, err)
	})
}

func TestUrlsRepo_GetFullUrl(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()
	const ttl = 1 * time.Second
	const ttlStr = "1 seconds"

	urlRepo := NewPgRepo(sqlxDB, ttl, ttlStr)

	t.Run("Success", func(t *testing.T) {
		originalUrl := "http://example.com"
		shortUrl := "mockRecord"

		rows := sqlmock.NewRows([]string{"original_url"}).AddRow(originalUrl)
		mock.ExpectQuery(GetFullUrlByShortQuery).WithArgs(shortUrl).WillReturnRows(rows)

		result, err := urlRepo.GetFullUrl(context.Background(), shortUrl)

		require.NoError(t, err)
		require.Equal(t, originalUrl, result)
	})
}
