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
	expiration := time.Now().Add(1 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	c.SetCookie(cookie)
	url := h.config.GoogleOAuth.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback is
func (h *Handler) GoogleCallback(c echo.Context) error {
	cookie, err := c.Cookie("oauthstate")

	if err != nil {
		log.Fatalln("SSIBAL ERROR")
	}

	state := c.FormValue("state")
	log.Println("================")
	log.Println(cookie.Value)
	log.Println(state)
	log.Println("================")
	// if c.FormValue("state") != oauthstate.Value {
	// 	log.Printf("invalid google oauth state cookie:%s state:%s\n", oauthstate.Value, c.FormValue("state"))
	// 	return c.Redirect(http.StatusTemporaryRedirect, "/error")
	// }

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

// NaverLogin is
func (h *Handler) NaverLogin(c echo.Context) error {
	return c.String(http.StatusCreated, "naver Login")
}

// NaverCallback is
func (h *Handler) NaverCallback(c echo.Context) error {
	return c.String(http.StatusOK, "call back")
}
