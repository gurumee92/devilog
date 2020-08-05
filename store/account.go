package store

import (
	"errors"

	"github.com/gurumee92/devilog/model"
	"github.com/jinzhu/gorm"
)

// AccountStore is
type AccountStore struct {
	db *gorm.DB
}

// Save a post
func (store *AccountStore) Save(account *model.Account) (*model.Account, error) {
	db := store.db
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(account).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

// FindAccounts is
func (store *AccountStore) FindAccounts(count, page int) ([]model.Account, error) {
	offset := (page - 1) * count
	db := store.db
	var accounts []model.Account
	err := db.Order("created_at desc").Limit(count).Offset(offset).Find(&accounts).Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindByID is..
func (store *AccountStore) FindByID(id int) (*model.Account, error) {
	var account model.Account
	db := store.db
	err := db.Find(&account, id).Error

	if err != nil {
		return nil, err
	}

	if account.ID == 0 {
		return nil, errors.New("Account isn't exist")
	}

	return &account, nil
}

// FindByEmail is
func (store *AccountStore) FindByEmail(email string) (*model.Account, error) {
	var account model.Account
	db := store.db
	err := db.Where("email = ?", email).First(&account).Error

	if err != nil {
		return nil, err
	}

	if account.ID == 0 {
		return nil, errors.New("Account isn't exist")
	}

	return &account, nil
}

// Update is
func (store *AccountStore) Update(account *model.Account) (*model.Account, error) {
	db := store.db
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(account).UpdateColumns(model.Account{
			Email:    account.Email,
			Password: account.Password,
			Username: account.Username,
			Picture:  account.Picture,
		}).Error

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

// DeleteByID is
func (store *AccountStore) DeleteByID(id int) error {
	db := store.db
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(model.Account{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// NewAccountStore is
func NewAccountStore(db *gorm.DB) *AccountStore {
	return &AccountStore{db: db}
}
