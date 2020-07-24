package config

import (
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Config is
type Config struct {
	ApplicationPath string
	DatabaseDialect string
	DatabaseURL     string
	IsProduct       bool
	GoogleOAuth     oauth2.Config
}

// GetConfig is
func GetConfig() *Config {
	scopesString := os.Getenv("GOGGLE_SCOPES")
	scopes := make([]string, 0)

	for _, s := range strings.Split(scopesString, ",") {
		scopes = append(scopes, "https://www.googleapis.com/auth/userinfo."+s)
	}

	googleOAuth := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       scopes,
	}
	return &Config{
		ApplicationPath: os.Getenv("APPLICATION_PATH"),
		DatabaseDialect: os.Getenv("DATABASE_DIALECT"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		IsProduct:       true,
		GoogleOAuth:     googleOAuth,
	}
}

// GetTestConfig is
func GetTestConfig() *Config {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	s := strings.Split(wd, "/")
	path := strings.Join(s[:len(s)-1], "/")
	googleOAuth := oauth2.Config{
		ClientID:     "test",
		ClientSecret: "test",
		Endpoint:     google.Endpoint,
		RedirectURL:  "",
		Scopes:       []string{},
	}
	return &Config{
		ApplicationPath: path,
		DatabaseDialect: "sqlite3",
		DatabaseURL:     "./test.db",
		IsProduct:       false,
		GoogleOAuth:     googleOAuth,
	}
}
