package comment

import (
	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/movie-service/usecase"
	"github.com/phamtrung99/movie-service/usecase/comment"
)

type Route struct {
	commentUseCase comment.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		commentUseCase: useCase.Comment,
	}

	group.GET("", r.GetList)
	group.POST("", r.Insert)
	group.DELETE("", r.Delete)
}
