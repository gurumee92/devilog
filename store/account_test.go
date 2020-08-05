package store

import (
	"strconv"
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

func TestAccountFindByIdSuccess(t *testing.T) {
	id := 1
	test := "test" + strconv.Itoa(id)
	account, err := accountStore.FindByID(id)

	assert.NoError(t, err)
	assert.Equal(t, uint(id), account.ID)
	assert.Equal(t, test, account.Email)
	assert.Equal(t, test, account.Password)
	assert.Equal(t, test, account.Username)
	assert.Equal(t, test, account.Picture)
}

func TestAccountFindByIdFailed(t *testing.T) {
	id := 10
	account, err := accountStore.FindByID(id)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestAccountFindByEmailSuccess(t *testing.T) {
	id := 1
	test := "test" + strconv.Itoa(id)
	account, err := accountStore.FindByEmail(test)

	assert.NoError(t, err)
	assert.Equal(t, uint(id), account.ID)
	assert.Equal(t, test, account.Email)
	assert.Equal(t, test, account.Password)
	assert.Equal(t, test, account.Username)
	assert.Equal(t, test, account.Picture)
}

func TestAccountFindByEmailFailed(t *testing.T) {
	id := 10
	test := "test" + strconv.Itoa(id)
	account, err := accountStore.FindByEmail(test)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestUpdateAccount(t *testing.T) {
	account, err := accountStore.FindByID(1)
	assert.NoError(t, err)

	account.Email = "Updated"
	account.Password = "Updated"
	account.Username = "Updated"
	account.Picture = "Updated"
	updated, err := accountStore.Update(account)
	assert.NoError(t, err)
	assert.Equal(t, account.Email, updated.Email)
	assert.Equal(t, account.Password, updated.Password)
	assert.Equal(t, account.Username, updated.Username)
	assert.Equal(t, account.Picture, updated.Picture)
}

func TestFindAccounts(t *testing.T) {
	tearDown()
	setup()
	accounts, err := accountStore.FindAccounts(5, 2)
	assert.NoError(t, err)
	assert.Equal(t, len(accounts), 0)

	accounts, err = accountStore.FindAccounts(5, 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(accounts))

	for idx, account := range accounts {
		assert.Equal(t, 2-idx, int(account.ID))
	}
}

func TestAccountDeleteByID(t *testing.T) {
	tearDown()
	setup()
	accounts, err := accountStore.FindAccounts(5, 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(accounts))

	id := 1
	err = accountStore.DeleteByID(id)
	assert.NoError(t, err)

	accounts, err = accountStore.FindAccounts(5, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(accounts))
}
