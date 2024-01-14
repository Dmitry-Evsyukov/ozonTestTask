package urlDelivery

import (
	"context"
	"github.com/labstack/echo/v4"
	"main/internal/models"
	"main/internal/url"
	httpError "main/pkg/err"
	"net/http"
)

type Handler struct {
	urlUseCase url.UseCase
}

func NewHandler(uuc url.UseCase) Handler {
	return Handler{urlUseCase: uuc}
}

func (h Handler) CreateShortUrl(c echo.Context) error {
	var fullUrl models.Url
	if err := c.Bind(&fullUrl); err != nil {
		return c.JSON(httpError.HandleError(err))
	}

	shortUrl, err := h.urlUseCase.CreateShortUrlAndSave(context.TODO(), fullUrl)
	if err != nil {
		c.JSON(httpError.HandleError(err))
	}

	return c.JSON(http.StatusOK, shortUrl)
}

func (h Handler) GetFullUrl(c echo.Context) error {
	var shortUrl models.Url
	shortUrl.Url = c.Param("url")

	fullUrl, err := h.urlUseCase.GetFullUrl(context.TODO(), shortUrl)
	if err != nil {
		return c.JSON(httpError.HandleError(err))
	}

	return c.JSON(http.StatusOK, fullUrl)
}
