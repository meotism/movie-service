package myerror

import (
	"net/http"

	"github.com/phamtrung99/gopkg/apperror"
)

// ErrFile.
func ErrFileOver5MB(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 40000010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "file is over 5mb",
		Message:   "file is over 5mb",
	}
}

func ErrOpenFile(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 40000011,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "fail to open file",
		Message:   "fail to open file",
	}
}

func ErrReadBufferFail(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 40000012,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "fail to read file buffer",
		Message:   "fail to read file buffer",
	}
}

func ErrNotImageFile(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 40000013,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "file ia not an image",
		Message:   "file ia not an image",
	}
}
