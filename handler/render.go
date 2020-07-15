package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// IndexPage function is
func (h *Handler) IndexPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// CreatePostPage is
func (h *Handler) CreatePostPage(c echo.Context) error {
	return c.Render(http.StatusOK, "create_post", nil)
}

// UpdatePostPage is
func (h *Handler) UpdatePostPage(c echo.Context) error {
	return c.Render(http.StatusOK, "update_post", nil)
}
