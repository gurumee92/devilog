package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndComparePassword(t *testing.T) {
	password := "1234"
	hash, err := Generate(password)

	assert.NoError(t, err)

	isMatch, err := Compare(hash, password)
	assert.NoError(t, err)
	assert.True(t, isMatch)
}
