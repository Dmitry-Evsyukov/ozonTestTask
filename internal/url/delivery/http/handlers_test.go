package urlDelivery

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"main/internal/models"
	"main/internal/url/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUrlHandlers_CreateShortUrl(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUrlUC := mock.NewMockUseCase(ctrl)
	urlHandler := NewHandler(mockUrlUC)

	originalUrl := models.Url{Url: "http://example.com"}
	shortUrl := models.Url{Url: "sgsdfsdf"}
	buf, err := json.Marshal(originalUrl)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/shorten_url", bytes.NewReader(buf))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(req, res)

	mockUrlUC.EXPECT().CreateShortUrlAndSave(context.Background(), originalUrl).Return(shortUrl, nil)

	err = urlHandler.CreateShortUrl(ctx)
	require.NoError(t, err)
}

func TestUrlHandlers_GetFullUrl(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUrlUC := mock.NewMockUseCase(ctrl)
	urlHandler := NewHandler(mockUrlUC)

	originalUrl := models.Url{Url: "http://example.com"}
	shortUrl := models.Url{Url: "sgsdfsdf"}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(req, res)
	ctx.SetParamNames("url")
	ctx.SetParamValues(shortUrl.Url)

	mockUrlUC.EXPECT().GetFullUrl(context.Background(), shortUrl).Return(originalUrl, nil)

	err := urlHandler.GetFullUrl(ctx)
	require.NoError(t, err)

	var response map[string]string
	_ = json.Unmarshal(res.Body.Bytes(), &response)
	require.Equal(t, originalUrl.Url, response["url"])
}
