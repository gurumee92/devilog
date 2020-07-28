package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GoogleLogin is
func (h *Handler) GoogleLogin(c echo.Context) error {
	return c.String(http.StatusCreated, "google Login")
}

// NaverLogin is
func (h *Handler) NaverLogin(c echo.Context) error {
	return c.String(http.StatusCreated, "naver Login")
}
