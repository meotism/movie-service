package user

import (
	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/movie-service/usecase"
	"github.com/phamtrung99/movie-service/usecase/user"
)

type Route struct {
	userUseCase user.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		userUseCase: useCase.User,
	}

	group.GET("/me", r.GetMe)
	group.PUT("/me", r.Update)
	group.POST("/changepassword", r.ChangePassword)
}
