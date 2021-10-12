package movie

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
	"github.com/phamtrung99/movie-service/model"
	"github.com/phamtrung99/movie-service/usecase/movie"
)

func (r *Route) SearchMovie(c echo.Context) error {
	var (
		ctx       = &utils.CustomEchoContext{Context: c}
		appError  = apperror.AppError{}
		paginator = model.Paginator{}
	)

	if err := c.Bind(&paginator); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	searchText := c.QueryParam("search")

	req := movie.ListMovieRequest{
		Paginator: &paginator,
	}
	// Bind order by
	if err := c.Bind(&req); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	res, err := r.movieUseCase.SearchMovie(ctx, searchText, req)

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, res)
}
