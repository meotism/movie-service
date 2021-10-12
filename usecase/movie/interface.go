package movie

import (
	"context"

	"github.com/phamtrung99/movie-service/model"
)

// IUsecase .
type IUsecase interface {
	SearchMovie(ctx context.Context, searchText string, req ListMovieRequest) (*model.MovieResult, error)
	Insert(ctx context.Context, req InsertMovieRequest) (*model.Movie, error)
	Delete(ctx context.Context, req DeleteMovieRequest) error
	Update(ctx context.Context, req UpdateMovieRequest) (*model.Movie, error)
}
