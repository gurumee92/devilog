package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// SocialLogin is ...
func (h *Handler) SocialLogin(c echo.Context) error {
	provider := c.Param("social_provider")
	state := generateOauthState()
	cookie := new(http.Cookie)
	cookie.Name = "state"
	cookie.Value = state
	cookie.Expires = time.Now().Add(1 * 24 * time.Hour)
	cookie.Path = "/login/oauth2/code/" + provider
	c.SetCookie(cookie)

	if provider == "google" {
		url := h.config.GoogleOAuth.AuthCodeURL(state)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	}

	url := h.config.NaverOAuth.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback is
func (h *Handler) GoogleCallback(c echo.Context) error {
	cookie, err := c.Cookie("state")

	if err != nil {
		return err
	}

	if c.FormValue("state") != cookie.Value {
		errMessage := fmt.Sprintf("invalid google oauth state cookie:%s state:%s\n", cookie.Value, c.FormValue("state"))
		return errors.New(errMessage)
	}

	code := c.FormValue("code")
	token, err := h.config.GoogleOAuth.Exchange(context.Background(), code)

	if err != nil {
		return err
	}

	if !token.Valid() {
		return errors.New("invalid token")
	}

	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?access_token=%v", token.AccessToken)
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	jsonMap := make(map[string]interface{})
	json.Unmarshal(contents, &jsonMap)
	log.Println(jsonMap)
	// id := jsonMap["id"]
	// email := jsonMap["email"]
	// username := jsonMap["name"]
	// picture := jsonMap["picture"]
	// log.Println(id, email, username, picture)
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

// NaverCallback is
func (h *Handler) NaverCallback(c echo.Context) error {
	cookie, err := c.Cookie("state")

	if err != nil {
		return err
	}

	if c.FormValue("state") != cookie.Value {
		errMessage := fmt.Sprintf("invalid naver oauth state cookie:%s state:%s\n", cookie.Value, c.FormValue("state"))
		return errors.New(errMessage)
	}

	code := c.FormValue("code")
	token, err := h.config.NaverOAuth.Exchange(context.Background(), code)

	if err != nil {
		return err
	}

	if !token.Valid() {
		return errors.New("invalid token")
	}

	url := "https://openapi.naver.com/v1/nid/me"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	jsonMap := make(map[string]interface{})
	json.Unmarshal(contents, &jsonMap)
	jsonResp := jsonMap["response"].(map[string]interface{})
	log.Println(jsonResp)
	// id := jsonResp["id"]
	// email := jsonResp["email"]
	// username := jsonResp["name"]
	// picture := jsonResp["profile_image"]
	// log.Println(id, email, username, picture)
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}