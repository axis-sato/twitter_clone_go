package handler

import (
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/null"

	"github.com/c8112002/twitter_clone_go/models"
	"github.com/c8112002/twitter_clone_go/router"
	"github.com/c8112002/twitter_clone_go/store"
	"github.com/c8112002/twitter_clone_go/utils"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/c8112002/twitter_clone_go/db"
	_ "github.com/go-sql-driver/mysql"
)

var (
	h   *Handler
	d   *sql.DB
	us  *store.UserStore
	ts  *store.TweetStore
	ctx context.Context
)

func setup() {
	var err error
	d, err = db.TestDB(false, utils.Location())

	if err != nil {
		panic(err.Error())
	}

	_ = router.New()

	ctx = context.Background()
	us = store.NewUserStore(d, ctx)
	ts = store.NewTweetStore(d, ctx)

	h = NewHandler(us, ts)

	if err := db.MigrateTestDB(d); err != nil {
		panic(err.Error())
	}

	if err := loadFixtures(); err != nil {
		panic(err.Error())
	}
}

func tearDown() {
	if err := db.DropTestDB(d); err != nil {
		panic(err.Error())
	}
}

func loadFixtures() error {
	users := []models.User{
		models.User{
			Name:    "鈴木 一郎",
			Icon:    "https://icon/1",
			Profile: "こんにちは。鈴木一郎です。",
		},
		models.User{
			Name:    "佐藤 二郎",
			Icon:    "https://icon/2",
			Profile: "こんにちは。佐藤二郎です。",
		},
		models.User{
			Name:    "田中 三郎",
			Icon:    "https://icon/3",
			Profile: "こんにちは。田中三郎です。",
		},
		models.User{
			Name:    "高橋 四郎",
			Icon:    "https://icon/4",
			Profile: "こんにちは。高橋四郎です。",
		},
		models.User{
			Name:      "橋本 五郎",
			Icon:      "https://icon/5",
			Profile:   "こんにちは。橋本五郎です。",
			DeletedAt: null.NewTime(time.Now(), true),
		},
	}

	tx, err := d.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, u := range users {
		if err := u.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
