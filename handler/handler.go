package handler

import (
	"github.com/gurumee92/devilog/config"
	"github.com/gurumee92/devilog/store"
)

// Handler is ...
type Handler struct {
	config    *config.Config
	postStore *store.PostStore
}

// NewHandler is Function, return Handler
func NewHandler(config *config.Config, postStore *store.PostStore) *Handler {
	return &Handler{
		config:    config,
		postStore: postStore,
	}
}
