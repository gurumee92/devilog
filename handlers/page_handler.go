package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// IndexPage is rendering Index Page
func IndexPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Echo")
}
