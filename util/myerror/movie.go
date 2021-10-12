package myerror

import (
	"net/http"

	"github.com/phamtrung99/gopkg/apperror"
)

// ErrAuthLogin .
func ErrGetMovie(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "fail to get movie",
		Message:   "fail to get movie",
	}
}

func ErrSearchTexFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000011,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "search text format is wrong",
		Message:   "search text format is wrong.",
	}
}

func ErrIsAdultFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000012,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "is_adult format is wrong",
		Message:   "is_adult format is wrong.",
	}
}

func ErrOriLanguageFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000013,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "original language format is wrong",
		Message:   "original language format is wrong.",
	}
}

func ErrOriTitleFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000014,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "original title format is wrong",
		Message:   "original title format is wrong.",
	}
}

func ErrOverviewFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000015,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "overview format is wrong",
		Message:   "overview format is wrong.",
	}
}

func ErrPopularityFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000015,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "popularity format is wrong",
		Message:   "popularity format is wrong.",
	}
}

func ErrReleaseDateFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000016,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "release date format is wrong",
		Message:   "release date format is wrong.",
	}
}

func ErrDurationFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000017,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "duration format is wrong",
		Message:   "duration format is wrong.",
	}
}

func ErrRatingAvgFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000018,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "rating everage format is wrong",
		Message:   "rating everage format is wrong.",
	}
}

func ErrStatusFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000019,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "status format is wrong",
		Message:   "status format is wrong.",
	}
}

func ErrSpokenLanguageFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000020,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "spoken language format is wrong",
		Message:   "spoken language format is wrong.",
	}
}


func ErrAgeFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000021,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "age format is wrong",
		Message:   "age format is wrong.",
	}
}
