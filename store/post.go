package store

import (
	"errors"

	"github.com/gurumee92/devilog/model"
	"github.com/jinzhu/gorm"
)

// PostStore is store posts
type PostStore struct {
	db *gorm.DB
}

// Save a post
func (store *PostStore) Save(post *model.Post) (*model.Post, error) {
	db := store.db
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(post).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

// FindByID find post by id
func (store *PostStore) FindByID(id int) (*model.Post, error) {
	var post model.Post
	db := store.db
	err := db.Find(&post, id).Error

	if err != nil {
		return nil, err
	}

	if post.ID == 0 {
		return nil, errors.New("User isn't exist")
	}

	return &post, nil
}

// Update is post update
func (store *PostStore) Update(post *model.Post) (*model.Post, error) {
	db := store.db
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(post).UpdateColumns(model.Post{
			Title:   post.Title,
			Content: post.Content,
			Author:  post.Author,
		}).Error

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

// FindPosts is
func (store *PostStore) FindPosts(count, page int) ([]model.Post, error) {
	offset := (page - 1) * count
	db := store.db
	var posts []model.Post
	err := db.Order("created_at desc").Limit(count).Offset(offset).Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

// DeleteByID is
func (store *PostStore) DeleteByID(id int) error {
	db := store.db
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(model.Post{}, "id == ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// NewPostStore is create instance PostStore
func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{db: db}
}
