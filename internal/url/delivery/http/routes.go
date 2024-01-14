package urlDelivery

import (
	"github.com/labstack/echo/v4"
	"main/internal/middleware"
)

func MapUrlRoutes(s *echo.Echo, h Handler, manager middleware.Manager) {
	s.POST("/shorten_url", h.CreateShortUrl, manager.RequestLoggerMiddleware)
	s.GET("/:url", h.GetFullUrl, manager.RequestLoggerMiddleware)
}
