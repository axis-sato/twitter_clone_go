package handler

import (
	"time"

	"github.com/c8112002/twitter_clone_go/entities"
)

type userResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	Profile       string `json:"profile"`
	IsFollower    bool   `json:"is_follower"`
	FolloweeCount int    `json:"followee_count"`
}

type usersResponse struct {
	Users             []*userResponse `json:"users"`
	ContainsFirstUser bool            `json:"contains_first_user"`
}

func newUserResponse(u *entities.User, isFollower bool) *userResponse {
	return &userResponse{
		ID:            u.ID,
		Name:          u.Name,
		Icon:          u.Icon,
		Profile:       u.Profile,
		IsFollower:    isFollower,
		FolloweeCount: len(u.Followees),
	}
}

func newEmptyUsersResponse() *usersResponse {
	return &usersResponse{
		Users:             []*userResponse{},
		ContainsFirstUser: false,
	}
}

type tweetResponse struct {
	ID        uint          `json:"id"`
	Tweet     string        `json:"tweet"`
	CreatedAt time.Time     `json:"created_at"`
	User      *userResponse `json:"user"`
	Like      int           `json:"like"`
	IsLiked   bool          `json:"is_liked"`
}

type tweetsResponse struct {
	Tweets             []*tweetResponse `json:"tweets"`
	ContainsFirstTweet bool             `json:"contains_first_tweet"`
}

func newTweetResponse(t *entities.Tweet, isLiked bool) *tweetResponse {
	return &tweetResponse{
		ID:        t.ID,
		Tweet:     t.Tweet,
		CreatedAt: t.CreatedAt,
		User:      newUserResponse(&t.User, t.User.IsFollowedBy(entities.LoginUserID)),
		Like:      t.Likes(),
		IsLiked:   isLiked,
	}
}

func newEmptyTweetsResponse() *tweetsResponse {
	return &tweetsResponse{
		Tweets:             []*tweetResponse{},
		ContainsFirstTweet: false,
	}
}
