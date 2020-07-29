package handler

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GoogleLogin is
func (h *Handler) GoogleLogin(c echo.Context) error {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := new(http.Cookie)
	cookie.Name = "oauthstate"
	cookie.Value = state
	cookie.Expires = time.Now().Add(1 * 24 * time.Hour)
	c.SetCookie(cookie)
	url := h.config.NaverOAuth.AuthCodeURL(state)
	return c.Redirect(http.StatusPermanentRedirect, url)
}

// GoogleCallback is
func (h *Handler) GoogleCallback(c echo.Context) error {
	cookie, err := c.Cookie("oauthstate")

	if err != nil {
		log.Println("error: ", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	if c.FormValue("state") != cookie.Value {
		log.Printf("invalid google oauth state cookie:%s state:%s\n", cookie.Value, c.FormValue("state"))
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

// NaverLogin is
func (h *Handler) NaverLogin(c echo.Context) error {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := new(http.Cookie)
	cookie.Name = "oauthstate"
	cookie.Value = state
	cookie.Expires = time.Now().Add(1 * 24 * time.Hour)
	c.SetCookie(cookie)
	url := h.config.NaverOAuth.AuthCodeURL(state)
	return c.Redirect(http.StatusPermanentRedirect, url)
}

// NaverCallback is
func (h *Handler) NaverCallback(c echo.Context) error {
	cookie, err := c.Cookie("oauthstate")

	if err != nil {
		log.Println("error: ", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	if c.FormValue("state") != cookie.Value {
		log.Printf("invalid google oauth state cookie:%s state:%s\n", cookie.Value, c.FormValue("state"))
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
