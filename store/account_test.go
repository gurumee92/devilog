package store

import (
	"testing"

	"github.com/gurumee92/devilog/model"
	"github.com/stretchr/testify/assert"
)

func TestSaveAccountSuccess(t *testing.T) {
	email := "test3"
	password := "test3"
	username := "test3"
	picture := "test3"
	account := model.Account{
		Email:    email,
		Password: password,
		Username: username,
		Picture:  picture,
	}
	saved, err := accountStore.Save(&account)

	assert.Equal(t, err, nil)
	assert.Less(t, uint(1), saved.ID)
	assert.Equal(t, email, saved.Email)
	assert.Equal(t, password, saved.Password)
	assert.Equal(t, username, saved.Username)
	assert.Equal(t, picture, saved.Picture)
}

func TestSaveAccountFail(t *testing.T) {
	email := "test1"
	password := ""
	username := ""
	picture := ""
	account := model.Account{
		Email:    email,
		Password: password,
		Username: username,
		Picture:  picture,
	}
	saved, err := accountStore.Save(&account)
	assert.Error(t, err)
	assert.Nil(t, saved)
}
