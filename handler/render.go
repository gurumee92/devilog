package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// IndexPage function is
func (h *Handler) IndexPage(c echo.Context) error {
	store := h.postStore
	posts, err := store.FindPosts(5, 0)

	if err != nil {
		c.Render(http.StatusInternalServerError, "error", err)
	}

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Posts": posts,
	})
}

// CreatePostPage is
func (h *Handler) CreatePostPage(c echo.Context) error {
	return c.Render(http.StatusOK, "create_post", nil)
}

// UpdatePostPage is
func (h *Handler) UpdatePostPage(c echo.Context) error {
	return c.Render(http.StatusOK, "update_post", nil)
}
