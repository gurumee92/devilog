package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, IndexPage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello Echo", rec.Body.String())
	}
}

func TestLoginPage(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, LoginPage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Login", rec.Body.String())
	}
}

func TestSignupPage(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/signup", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, SignupPage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Signup", rec.Body.String())
	}
}

func testUserHomePage(t *testing.T, username string) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/home/:username", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/home/:username")
	c.SetParamNames("username")
	c.SetParamValues(username)

	// Assertions
	if assert.NoError(t, UserHomePage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, username, rec.Body.String())
	}
}

func TestUserHomePage(t *testing.T) {
	usernames := []string{"test1", "test2"}

	for _, username := range usernames {
		testUserHomePage(t, username)
	}
}
