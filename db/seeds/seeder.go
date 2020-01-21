package main

import (
	"context"
	db2 "github.com/c8112002/twitter_clone_go/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db2.New()

	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err.Error())
		}
	}()

	ctx := context.Background()
	makeDummyUsers(ctx, db)
	makeDummyTweets(ctx, db)
	makeDummyFollowers(ctx, db)
	makeDummyLikes(ctx, db)
}
