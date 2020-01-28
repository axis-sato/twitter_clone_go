package handler

import "github.com/c8112002/twitter_clone_go/entities"

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
