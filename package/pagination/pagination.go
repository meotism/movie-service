package pagination

import (
	"strconv"

	checkform "github.com/phamtrung99/movie-service/package/checkForm"

	"github.com/labstack/echo/v4"
)

func GetPageQueryParam(c echo.Context) (bool, string, int) {
	if c.QueryParam("page") != "" {
		isPage, msg := checkform.CheckFormatValue("page", c.QueryParam("page"))
		if !isPage {
			return false, msg, 1
		}
		page, _ := strconv.Atoi(msg)
		return true, "", page
	}
	return true, "", 1
}

func CountTotalPage(totalResult int64, rowPerPage int64) int64 {

	if totalResult%rowPerPage > 0 {
		return (totalResult / rowPerPage) + 1
	}

	return totalResult / rowPerPage
}
