package handler

import (
	"github.com/gurumee92/devilog/config"
	"github.com/gurumee92/devilog/store"
)

// Handler is ...
type Handler struct {
	config       *config.Config
	postStore    *store.PostStore
	accountStore *store.AccountStore
}

// NewHandler is Function, return Handler
func NewHandler(config *config.Config, postStore *store.PostStore, accountStore *store.AccountStore) *Handler {
	return &Handler{
		config:       config,
		postStore:    postStore,
		accountStore: accountStore,
	}
}
