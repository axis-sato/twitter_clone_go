package store

import (
	"context"
	"database/sql"

	"github.com/c8112002/twitter_clone_go/entities"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/c8112002/twitter_clone_go/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) FetchUsers(lastID int, limit int) (entities.Users, error) {

	ctx := context.Background()

	ul, err := models.Users(
		qm.Where("id <= ?", lastID),
		qm.Where("deleted_at is null"),
		qm.Limit(limit),
		qm.OrderBy(models.UserColumns.ID+" desc"),
		qm.Load(qm.Rels(models.UserRels.Followers, models.FollowRels.Followee), qm.Where("deleted_at is null")),
		qm.Load(qm.Rels(models.UserRels.Followees, models.FollowRels.Follower), qm.Where("deleted_at is null")),
	).All(ctx, us.db)

	users := entities.Users{}

	for _, u := range ul {
		followers := entities.Followers{}
		followees := entities.Followees{}

		frl := u.R.Followers
		for _, f := range frl {
			fu := f.R.Followee
			if fu == nil {
				continue
			}
			fr := entities.NewFollower(fu.ID, fu.Name, fu.Icon, fu.Profile)
			followers = append(followers, fr)
		}

		fel := u.R.Followees
		for _, f := range fel {
			fu := f.R.Follower
			if fu == nil {
				continue
			}
			fe := entities.NewFollowee(fu.ID, fu.Name, fu.Icon, fu.Profile)
			followees = append(followees, fe)
		}

		user := entities.NewUser(u.ID, u.Name, u.Icon, u.Profile, followers, followees)
		users = append(users, user)
	}

	return users, err
}
