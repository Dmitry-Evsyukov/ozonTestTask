package urlUseCase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"main/internal/models"
	"main/internal/url/mock"
	"testing"
)

func TestUrlUC_SaveShortToFullUrl_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUrlRepo := mock.NewMockRepository(ctrl)
	urlUC := New(mockUrlRepo)

	originalUrl := models.Url{Url: "https://example.com"}

	mockUrlRepo.EXPECT().SaveShortToFullUrl(context.Background(), gomock.Any(), originalUrl.Url).Return(nil)

	shortUrl, err := urlUC.CreateShortUrlAndSave(context.Background(), originalUrl)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotEqual(t, "", shortUrl.Url)
}

func TestUrlUC_SaveShortToFullUrl_BadRequest(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUrlRepo := mock.NewMockRepository(ctrl)
	urlUC := New(mockUrlRepo)

	originalUrl := models.Url{Url: "example.com"}

	shortUrl, err := urlUC.CreateShortUrlAndSave(context.Background(), originalUrl)
	require.Error(t, err)
	require.Equal(t, "", shortUrl.Url)
}

func TestUrlUC_GetFullUrl_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUrlRepo := mock.NewMockRepository(ctrl)
	urlUC := New(mockUrlRepo)

	shortUrl := models.Url{Url: "short_url"}
	originalUrl := "http://example.com"

	mockUrlRepo.EXPECT().GetFullUrl(context.Background(), shortUrl.Url).Return(originalUrl, nil)

	originalUrlFromUC, err := urlUC.GetFullUrl(context.Background(), shortUrl)
	require.NoError(t, err)
	require.Equal(t, originalUrl, originalUrlFromUC.Url)
}
