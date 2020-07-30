package handler

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func generateOauthState() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}

// GoogleLogin is
func (h *Handler) GoogleLogin(c echo.Context) error {
	state := generateOauthState()
	cookie := new(http.Cookie)
	cookie.Name = "state"
	cookie.Value = state
	cookie.Expires = time.Now().Add(1 * 24 * time.Hour)
	cookie.Path = "/login/oauth2/code/google"
	c.SetCookie(cookie)
	url := h.config.GoogleOAuth.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback is
func (h *Handler) GoogleCallback(c echo.Context) error {
	cookie, err := c.Cookie("state")

	if err != nil {
		return err
	}

	if c.FormValue("state") != cookie.Value {
		log.Printf("invalid google oauth state cookie:%s state:%s\n", cookie.Value, c.FormValue("state"))
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

// NaverLogin is
func (h *Handler) NaverLogin(c echo.Context) error {
	state := generateOauthState()
	cookie := new(http.Cookie)
	cookie.Name = "state"
	cookie.Value = state
	cookie.Expires = time.Now().Add(1 * 24 * time.Hour)
	cookie.Path = "/login/oauth2/code/naver"
	c.SetCookie(cookie)
	url := h.config.NaverOAuth.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// NaverCallback is
func (h *Handler) NaverCallback(c echo.Context) error {
	cookie, err := c.Cookie("state")

	if err != nil {
		log.Println("error: ", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	if c.FormValue("state") != cookie.Value {
		log.Printf("invalid naver oauth state cookie:%s state:%s\n", cookie.Value, c.FormValue("state"))
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
