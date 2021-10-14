package movie

import (
	"context"
	"fmt"

	"github.com/phamtrung99/movie-service/model"
	checkform "github.com/phamtrung99/movie-service/package/checkForm"
	"github.com/phamtrung99/movie-service/util/myerror"
)

// ListRequest .
type ListMovieRequest struct {
	Filter    *model.MovieFilter
	Paginator *model.Paginator
	OrderBy   string `json:"order_by,omitempty" query:"order_by"`
	OrderType string `json:"order_type,omitempty" query:"order_type"`
}

func (u *Usecase) SearchMovie(ctx context.Context, searchText string, req ListMovieRequest) (*model.MovieResult, error) {

	//log.Fatal(req.Paginator)
	isTrue, strSearch := checkform.CheckFormatValue("search", searchText)
	if !isTrue {
		return nil, myerror.ErrSearchTexFormat(nil)
	}

	orders := make([]string, 0)
	if req.OrderBy != "" {
		orders = []string{fmt.Sprintf("%s %s", req.OrderBy, req.OrderType)}
	}

	paginator := &model.Paginator{
		Page:  1,
		Limit: 20,
	}

	if req.Paginator != nil {
		paginator = req.Paginator
	}

	movieList, err := u.movieRepo.SearchMovie(ctx, paginator, strSearch, req.Filter, orders)

	if err != nil {
		return nil, myerror.ErrGetMovie(err)
	}

	return movieList, nil
}
