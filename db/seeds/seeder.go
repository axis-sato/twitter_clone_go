package main

import (
	"context"

	"github.com/c8112002/twitter_clone_go/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	d, err := db.New(false)

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
