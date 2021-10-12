package movie

import (
	"context"

	"github.com/phamtrung99/movie-service/model"
)

// Repository .
type Repository interface {
	Find(ctx context.Context,
		conditions []model.Condition,
		paginator *model.Paginator,
		orders []string,
	) (*model.MovieResult, error)
	SearchMovie(
		ctx context.Context,
		paginator *model.Paginator,
		str string,
		filter *model.MovieFilter,
		orders []string) (*model.MovieResult, error)
	GetById(ctx context.Context, ID int64) (*model.Movie, error)
	Insert(ctx context.Context, movie *model.Movie) (*model.Movie, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, movie *model.Movie) (*model.Movie, error)
}
