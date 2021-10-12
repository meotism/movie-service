package usecase

import (
	"github.com/phamtrung99/movie-service/repository"
	"github.com/phamtrung99/movie-service/usecase/movie"
)

type UseCase struct {
	Movie movie.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		Movie: movie.New(repo),
	}
}
