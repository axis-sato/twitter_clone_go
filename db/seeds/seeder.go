package main

import (
	"context"

	"github.com/c8112002/twitter_clone_go/db"
	"github.com/c8112002/twitter_clone_go/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loc := utils.Location()
	d, err := db.New(false, loc)

	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err := d.Close(); err != nil {
			panic(err.Error())
		}
	}()

	ctx := context.Background()
	makeDummyUsers(ctx, d)
	makeDummyTweets(ctx, d)
	makeDummyFollows(ctx, d)
	makeDummyLikes(ctx, d)
}
