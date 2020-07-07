package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// IndexPage is rendering Index Page
func IndexPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Echo")
}

// LoginPage is rendering Login Page
func LoginPage(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}

// SignupPage is rendering Signup Page
func SignupPage(c echo.Context) error {
	return c.String(http.StatusOK, "Signup")
}

// UserHomePage is rendering User's Home Page
func UserHomePage(c echo.Context) error {
	username := c.Param("username")
	return c.String(http.StatusOK, username)
}
