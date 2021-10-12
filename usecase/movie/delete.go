package movie

import (
	"context"

	"github.com/phamtrung99/movie-service/util/myerror"
)

type DeleteMovieRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) Delete(ctx context.Context, req DeleteMovieRequest) error {

	movie, err := u.movieRepo.GetById(ctx, req.ID)

	if err != nil {
		return myerror.ErrGetMovie(err)
	}

	err = u.movieRepo.Delete(ctx, movie.ID)

	return err
}
