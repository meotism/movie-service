package imgvalid

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/phamtrung99/movie-service/util/myerror"
)

func CheckImage(file *multipart.FileHeader) (string, error) {

	if file.Size > 5242880 {
		return "", myerror.ErrFileOver5MB(nil)
	}

	src, err := file.Open()

	if err != nil {
		return "", myerror.ErrOpenFile(nil)
	}

	buff := make([]byte, 512)
	_, err = src.Read(buff)

	if err != nil {
		return "", myerror.ErrReadBufferFail(nil)
	}

	filetype := http.DetectContentType(buff)

	switch filetype {
	case "image/jpeg", "image/jpg", "image/gif", "image/png":

	default:
		return "", myerror.ErrNotImageFile(nil)
	}
	defer src.Close()

	return filetype, nil
}

func CopyFile(file *multipart.FileHeader, dstFileName string, path string) error {
	localFile, err := os.Create(path + dstFileName)

	if err != nil {
		return err
	}

	defer localFile.Close()

	rootFile, err := file.Open()

	if err != nil {
		return err
	}

	if _, err := io.Copy(localFile, rootFile); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return err
	}

	return nil
}
