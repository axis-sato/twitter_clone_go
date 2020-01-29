package store

import (
	"github.com/c8112002/twitter_clone_go/entities"
	"github.com/c8112002/twitter_clone_go/models"
)

func createUser(u models.User) *entities.User {
	followers := createFollowersFrom(u)
	followees := createFolloweesFrom(u)

	return entities.NewUser(u.ID, u.Name, u.Icon, u.Profile, *followers, *followees)
}

func createFollowersFrom(u models.User) *entities.Followers {
	followers := entities.Followers{}
	for _, f := range u.R.Followers {
		fu := f.R.Followee
		if fu == nil {
			continue
		}
		fr := entities.NewFollower(fu.ID, fu.Name, fu.Icon, fu.Profile)
		followers = append(followers, fr)
	}

	return &followers
}

func createFolloweesFrom(u models.User) *entities.Followees {
	followees := entities.Followees{}
	for _, f := range u.R.Followees {
		fu := f.R.Follower
		if fu == nil {
			continue
		}
		fe := entities.NewFollowee(fu.ID, fu.Name, fu.Icon, fu.Profile)
		followees = append(followees, fe)
	}

	return &followees
}

func createTweet(t models.Tweet) *entities.Tweet {
	u := t.R.User
	user := createUser(*u)
	likes := t.R.Likes
	var lusers entities.Users
	for _, like := range likes {
		luser := createUser(*like.R.User)
		lusers = append(lusers, luser)
	}
	return entities.NewTweet(t.ID, t.Tweet, t.CreatedAt, *user, lusers)
}
