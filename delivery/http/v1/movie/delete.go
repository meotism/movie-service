package movie

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
	"github.com/phamtrung99/movie-service/usecase/movie"
)

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	movieID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	req := movie.DeleteMovieRequest{
		ID: movieID,
	}

	err := r.movieUseCase.Delete(ctx, req)

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, nil)
}
