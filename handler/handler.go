package handler

import "github.com/gurumee92/devilog/store"

// Handler is ...
type Handler struct {
	postStore *store.PostStore
}

// NewHandler is Function, return Handler
func NewHandler(postStore *store.PostStore) *Handler {
	return &Handler{
		postStore: postStore,
	}
}
