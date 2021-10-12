package movie

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
	"github.com/phamtrung99/movie-service/usecase/movie"
)

func (r *Route) Update(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	form, err := c.MultipartForm()
	if err != nil {
		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	res, err := r.movieUseCase.Update(ctx, movie.UpdateMovieRequest{FormData: form})

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(c, res)
}
