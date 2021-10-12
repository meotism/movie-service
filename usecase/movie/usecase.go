package movie

import (
	"github.com/phamtrung99/movie-service/repository"
	"github.com/phamtrung99/movie-service/repository/movie"
)

type Usecase struct {
	movieRepo movie.Repository
}

// New .
func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		movieRepo: repo.Movie,
	}
}
