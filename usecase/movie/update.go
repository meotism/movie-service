package movie

import (
	"context"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/phamtrung99/movie-service/model"
	checkform "github.com/phamtrung99/movie-service/package/checkForm"
	imgvalid "github.com/phamtrung99/movie-service/package/fileValid"
	"github.com/phamtrung99/movie-service/util/myerror"
)

type UpdateMovieRequest struct {
	FormData *multipart.Form
	MovieID  int64 `json:"movie_id"` //get from Param
}

func (u *Usecase) Update(ctx context.Context, req UpdateMovieRequest) (*model.Movie, error) {

	//Get current movie info by movie id
	movie, err := u.movieRepo.GetById(ctx, req.MovieID)

	if err != nil {
		return nil, myerror.ErrGetMovie(err)
	}

	//Parse from multipart form to movie which need update

	//Check is_adult format
	if req.FormData.Value["is_adult"][0] != "" {
		isAdult, err := strconv.Atoi(req.FormData.Value["is_adult"][0])

		if err != nil || isAdult > 2 || isAdult < 0 {
			return nil, myerror.ErrIsAdultFormat(nil)
		}

		movie.IsAdult = isAdult
	}

	//Check original_language format
	if req.FormData.Value["original_language"][0] != "" {
		isTrue, msg := checkform.CheckFormatValue("original_language", req.FormData.Value["original_language"][0])

		if !isTrue {
			return nil, myerror.ErrOriLanguageFormat(nil)
		}

		movie.OriginalLanguage = msg
	}

	//Check original_title format
	if req.FormData.Value["original_title"][0] != "" {
		isTrue, msg := checkform.CheckFormatValue("original_title", req.FormData.Value["original_title"][0])

		if !isTrue {
			return nil, myerror.ErrOriTitleFormat(nil)
		}

		movie.OriginalTitle = msg
	}

	//Check overview format
	if req.FormData.Value["overview"][0] != "" {
		isTrue, msg := checkform.CheckFormatValue("overview", req.FormData.Value["overview"][0])

		if !isTrue {
			return nil, myerror.ErrOverviewFormat(nil)
		}

		movie.Overview = msg
	}

	//Check popularity format
	if req.FormData.Value["popularity"][0] != "" {
		popularity, err := strconv.ParseFloat(req.FormData.Value["popularity"][0], 64)

		if err != nil {
			return nil, myerror.ErrPopularityFormat(nil)
		}

		movie.Popularity = popularity
	}

	//Check movie link
	if req.FormData.Value["movie_link"][0] != "" {
		movie.MovieLink = req.FormData.Value["movie_link"][0]
	}

	//Check release date format
	if req.FormData.Value["release_date"][0] != "" {
		t, err := time.Parse(time.RFC3339, req.FormData.Value["release_date"][0])
		if err != nil {
			return nil, myerror.ErrReleaseDateFormat(nil)
		}

		movie.ReleaseDate = t
	}

	//Check duration format
	if req.FormData.Value["duration"][0] != "" {
		duration, err := strconv.Atoi(req.FormData.Value["duration"][0])

		if err != nil || duration < 0 {
			return nil, myerror.ErrDurationFormat(nil)
		}

		movie.Duration = duration
	}

	//Check spoken language format
	if req.FormData.Value["spoken_language"][0] != "" {
		isTrue, msg := checkform.CheckFormatValue("spoken_language", req.FormData.Value["spoken_language"][0])

		if !isTrue {
			return nil, myerror.ErrSpokenLanguageFormat(nil)
		}

		movie.SpokenLanguage = msg
	}

	//Check rating average format
	if req.FormData.Value["rating_average"][0] != "" {
		ratingAvg, err := strconv.ParseFloat(req.FormData.Value["rating_average"][0], 64)

		if err != nil {
			return nil, myerror.ErrRatingAvgFormat(nil)
		}

		movie.RatingAverage = ratingAvg
	}

	//Check status format
	if req.FormData.Value["status"][0] != "" {
		status, err := strconv.Atoi(req.FormData.Value["status"][0])

		if err != nil || status > 2 || status < 0 {
			return nil, myerror.ErrStatusFormat(nil)
		}

		movie.Status = status
	}

	if len(req.FormData.File["image"]) != 0 {
		file := req.FormData.File["image"][0]
		var imgFileName = "blank.png"
		var pathFile = "/public/cover/"

		filetype, err := imgvalid.CheckImage(file)

		if err != nil {
			return nil, err
		}

		imgFileName = strconv.FormatInt(movie.ID, 10) + "." + strings.Split(filetype, "/")[1]

		err = imgvalid.CopyFile(file, imgFileName, "."+pathFile)

		if err != nil {
			log.Fatal(err)
		}

		movie.Image = pathFile + imgFileName
	}

	result, err := u.movieRepo.Update(ctx, movie)

	if err != nil {
		return nil, err
	}

	return result, nil
}
