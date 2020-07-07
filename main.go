package main

import (
	"github.com/gurumee92/devilog/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.IndexPage)
	e.GET("/login", handlers.LoginPage)
	e.GET("/signup", handlers.SignupPage)
	e.GET("/home/:username", handlers.UserHomePage)
	e.Logger.Fatal(e.Start(":1323"))
}
