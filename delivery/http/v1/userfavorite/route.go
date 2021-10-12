package userfavorite

import (
	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/movie-service/usecase"
	"github.com/phamtrung99/movie-service/usecase/userfavorite"
)

type Route struct {
	userFavorUseCase userfavorite.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		userFavorUseCase: useCase.UserFavor,
	}

	group.GET("", r.GetFavorite)
	group.POST("", r.AddFavorite)
}
