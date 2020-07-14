package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestConifguration(t *testing.T) {
	config := GetTestConfig()
	// your app path
	// appPath := "/Users/gurumee/Workspaces/devilog"
	// assert.Equal(t, appPath, config.ApplicationPath)

	assert.Equal(t, "sqlite3", config.DatabaseDialect)
	assert.Equal(t, "./test.db", config.DatabaseURL)
}
