package handler

import (
	"os"
	"testing"

	"github.com/gurumee92/devilog/router"
	"github.com/labstack/echo/v4"
)

var (
	h *Handler
	e *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	e = router.NewRouter()
	h = NewHandler()
	h.Register(e)
}

func tearDown() {
	h = nil
	e = nil
}
