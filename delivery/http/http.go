package http

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/movie-service/config"
	"github.com/phamtrung99/movie-service/delivery/http/v1/movie"
	"github.com/phamtrung99/movie-service/repository"
	"github.com/phamtrung99/movie-service/usecase"
)

// NewHTTPHandler .
func NewHTTPHandler(repo *repository.Repository, ucase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	cfg := config.GetConfig()

	skipper := func(c echo.Context) bool {
		p := c.Request().URL.Path

		return strings.Contains(p, "/health_check")
	}

	e.Use(middleware.Auth(cfg.Jwt.Key, skipper, false))

	apiV1 := e.Group("/v1")
	movie.Init(apiV1.Group("/movies"), ucase)

	return e
}
