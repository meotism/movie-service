package movie

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
	"github.com/phamtrung99/movie-service/model"
	"github.com/phamtrung99/movie-service/usecase/movie"
)

func (r *Route) SearchMovie(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	isAdult, _ := strconv.Atoi(c.QueryParam("is_adult"))
	minRating, _ := strconv.Atoi(c.QueryParam("min_rating"))
	actorID, _ := strconv.ParseInt(c.QueryParam("actor_id"), 10, 64)
	cateID, _ := strconv.ParseInt(c.QueryParam("cate_id"), 10, 64)
	searchText := c.QueryParam("search")

	req := movie.ListMovieRequest{
		Paginator: &model.Paginator{
			Page:  page,
			Limit: limit,
		},
		Filter: &model.MovieFilter{
			IsAdult:   isAdult,
			ActorID:   actorID,
			CateID:    cateID,
			MinRating: minRating,
		},
	}

	res, err := r.movieUseCase.SearchMovie(ctx, searchText, req)

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, res)
}
