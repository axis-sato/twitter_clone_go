package handler

import "github.com/c8112002/twitter_clone_go/store"

type Handler struct {
	userStore  *store.UserStore
	tweetStore *store.TweetStore
}

func NewHandler(us *store.UserStore, ts *store.TweetStore) *Handler {
	return &Handler{userStore: us, tweetStore: ts}
}
