package urlRepository

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUrlsRepo_SaveShortToFull(t *testing.T) {
	const ttl = 300 * time.Millisecond
	urlRepo := NewInMemory(ttl)

	t.Run("Success", func(t *testing.T) {
		originalUrl := "http://example.com"
		shortUrl := "mockRecord"

		err := urlRepo.SaveShortToFullUrl(context.Background(), shortUrl, originalUrl)

		require.NoError(t, err)
		require.NotNil(t, urlRepo.shortToFull[shortUrl])
		require.Equal(t, urlRepo.shortToFull[shortUrl].url, originalUrl)

		time.Sleep(2 * ttl)
		require.NotEqual(t, urlRepo.shortToFull[shortUrl].url, originalUrl)
	})
	urlRepo.stopCleanUp()
}

func TestUrlsRepo_GetFullUrl(t *testing.T) {
	const ttl = 300 * time.Millisecond
	urlRepo := NewInMemory(ttl)

	t.Run("Success", func(t *testing.T) {
		originalUrl := "http://example.com"
		shortUrl := "mockRecord"
		urlRepo.shortToFull[shortUrl] = internal{
			timeCreation: time.Now(),
			url:          originalUrl,
		}

		fullUrlFromDb, err := urlRepo.GetFullUrl(context.Background(), shortUrl)

		require.NoError(t, err)
		require.NotNil(t, fullUrlFromDb)
		require.Equal(t, originalUrl, fullUrlFromDb)

		time.Sleep(2 * ttl)

		fullUrlFromDb, err = urlRepo.GetFullUrl(context.Background(), shortUrl)
		require.Error(t, err)
		require.Equal(t, "", fullUrlFromDb)
	})
	urlRepo.stopCleanUp()
}
