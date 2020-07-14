package config

import (
	"log"
	"os"
	"strings"
)

// Config is
type Config struct {
	ApplicationPath string
	DatabaseDialect string
	DatabaseURL     string
	IsProduct       bool
}

// GetConfig is
func GetConfig() Config {
	return Config{
		ApplicationPath: os.Getenv("APPLICATION_PATH"),
		DatabaseDialect: os.Getenv("DATABASE_DIALECT"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		IsProduct:       true,
	}
}

// GetTestConfig is
func GetTestConfig() Config {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	s := strings.Split(wd, "/")
	path := strings.Join(s[:len(s)-1], "/")

	return Config{
		ApplicationPath: path,
		DatabaseDialect: "sqlite3",
		DatabaseURL:     "./test.db",
		IsProduct:       false,
	}
}
