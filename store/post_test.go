package store

import (
	"strconv"
	"testing"

	"github.com/gurumee92/devilog/model"
	"github.com/stretchr/testify/assert"
)

func TestSavePost(t *testing.T) {
	title := "test title"
	content := "test content"
	author := "test author"
	post := model.Post{
		Title:   title,
		Content: content,
		Author:  author,
	}
	saved, err := postStore.Save(&post)

	assert.Equal(t, err, nil)
	assert.Less(t, uint(1), saved.ID)
	assert.Equal(t, title, saved.Title)
	assert.Equal(t, content, saved.Content)
	assert.Equal(t, author, saved.Author)
}

func TestFindByIdSuccess(t *testing.T) {
	id := 1
	test := "test" + strconv.Itoa(id)
	post, err := postStore.FindByID(id)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, uint(id), post.ID)
	assert.Equal(t, test, post.Title)
	assert.Equal(t, test, post.Content)
	assert.Equal(t, test, post.Author)
}

func TestFindByIdFailed(t *testing.T) {
	id := 10
	post, err := postStore.FindByID(id)

	assert.Equal(t, true, err != nil)
	assert.Equal(t, true, post == nil)
}

func TestUpdate(t *testing.T) {
	post, _ := postStore.FindByID(1)
	post.Title = "UpdatedTitle"
	post.Content = "UpdatedContent"
	post.Author = "UpdatedAuthor"
	updated, err := postStore.Update(post)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, post.Title, updated.Title)
	assert.Equal(t, post.Content, updated.Content)
	assert.Equal(t, post.Author, updated.Author)
}
