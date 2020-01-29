package store

import (
	"context"
	"database/sql"

	"github.com/c8112002/twitter_clone_go/entities"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/c8112002/twitter_clone_go/models"
)

type UserStore struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserStore(db *sql.DB, ctx context.Context) *UserStore {
	return &UserStore{db: db, ctx: ctx}
}

func (us *UserStore) FetchUsers(lastID int, limit int) (*entities.Users, error) {

	ul, err := models.Users(
		qm.Where("id <= ?", lastID),
		qm.Where("deleted_at is null"),
		qm.Limit(limit),
		qm.OrderBy(models.UserColumns.ID+" desc"),
		qm.Load(qm.Rels(models.UserRels.Followers, models.FollowRels.Followee), qm.Where("deleted_at is null")),
		qm.Load(qm.Rels(models.UserRels.Followees, models.FollowRels.Follower), qm.Where("deleted_at is null")),
	).All(us.ctx, us.db)

	users := entities.Users{}

	for _, u := range ul {
		user := createUser(*u)
		users = append(users, user)
	}

	return &users, err
}

func (us *UserStore) FetchFirstUser() (*entities.User, error) {
	u, err := models.Users(
		qm.Where("deleted_at is null"),
		qm.OrderBy(models.UserColumns.ID),
		qm.Limit(1),
		qm.Load(qm.Rels(models.UserRels.Followers, models.FollowRels.Followee), qm.Where("deleted_at is null")),
		qm.Load(qm.Rels(models.UserRels.Followees, models.FollowRels.Follower), qm.Where("deleted_at is null")),
	).One(us.ctx, us.db)

	if err != nil {
		return nil, err
	}

	return createUser(*u), err
}
