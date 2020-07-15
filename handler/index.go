package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index function is 
func (h *Handler) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "World")
}
