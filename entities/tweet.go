package entities

import "time"

type Tweet struct {
	ID        uint
	Tweet     string
	CreatedAt time.Time
	User      User
	LikeUsers Users
}

type Tweets = []*Tweet

func NewTweet(id uint, tweet string, createdAt time.Time, user User, likeUsers Users) *Tweet {
	return &Tweet{
		ID:        id,
		Tweet:     tweet,
		CreatedAt: createdAt,
		User:      user,
		LikeUsers: likeUsers,
	}
}

func (u *Tweet) Likes() int {
	return len(u.LikeUsers)
}

// 引数のuserIDにLikeされている場合はtrueを返す
func (u *Tweet) IsLikedBy(userID uint) bool {
	for _, u := range u.LikeUsers {
		if u.ID == userID {
			return true
		}
	}

	return false
}
