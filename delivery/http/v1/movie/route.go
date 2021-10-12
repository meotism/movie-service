package movie

import (
	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/movie-service/usecase"
	"github.com/phamtrung99/movie-service/usecase/movie"
)

type Route struct {
	movieUseCase movie.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		movieUseCase: useCase.Movie,
	}

	group.GET("", r.SearchMovie)
	group.POST("", r.Insert)
	group.DELETE("", r.Delete)
	group.PUT("", r.Update)
}
