package handler

import (
	"github.com/labstack/echo/v4"
)

// Register is routes function
func (h *Handler) Register(e *echo.Echo) {
	e.GET("/hello", h.Hello)
}
