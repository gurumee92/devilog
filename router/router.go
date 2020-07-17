package router

import (
	"github.com/gurumee92/devilog/config"
	custom "github.com/gurumee92/devilog/router/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// NewRouter is return customized Echo instance
func NewRouter(c *config.Config) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Static(c.ApplicationPath + "/public/static"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Validator = NewValidator()
	e.Renderer = custom.NewTemplate(c)
	return e
}
