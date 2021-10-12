package movie

import (
	"context"
	"time"

	"github.com/phamtrung99/movie-service/model"
)

type InsertMovieRequest struct {
	IsAdult          int       `json:"is_adult"`
	OriginalLanguage string    `json:"original_language"`
	OriginalTitle    string    `json:"original_title"`
	Overview         string    `json:"overview"`
	Popularity       float64   `json:"popularity"`
	MovieLink        string    `json:"movie_link"`
	ReleaseDate      time.Time `json:"release_date"`
	Duration         int       `json:"duration"`
	SpokenLanguage   string    `json:"spoken_language"`
	RatingAverage    float64   `json:"rating_average"`
	Status           int       `json:"status"`
}

func (u *Usecase) Insert(ctx context.Context, req InsertMovieRequest) (*model.Movie, error) {

	movie := &model.Movie{
		IsAdult:          req.IsAdult,
		Image:            "https://media.istockphoto.com/photos/popcorn-and-clapperboard-picture-id1191001701?s=612x612",
		OriginalLanguage: req.OriginalLanguage,
		OriginalTitle:    req.OriginalTitle,
		Overview:         req.Overview,
		Popularity:       req.Popularity,
		MovieLink:        req.MovieLink,
		ReleaseDate:      req.ReleaseDate,
		Duration:         req.Duration,
		SpokenLanguage:   req.SpokenLanguage,
		RatingAverage:    req.RatingAverage,
		Status:           req.Status,
	}

	result, err := u.movieRepo.Insert(ctx, movie)

	if err != nil {
		return nil, err
	}

	return result, nil
}
