package handler

import "github.com/c8112002/twitter_clone_go/store"

type Handler struct {
	userStore *store.UserStore
}

func NewHandler(us *store.UserStore) *Handler {
	return &Handler{userStore: us}
}
