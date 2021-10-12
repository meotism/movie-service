package repository

import (
	"context"

	"github.com/phamtrung99/movie-service/repository/movie"
	"gorm.io/gorm"
)

type Repository struct {
	Movie movie.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		Movie: movie.NewPGRepository(getSQLClient),
	}
}
