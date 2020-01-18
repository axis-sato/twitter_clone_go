package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "twitter_clone:twitter_clone_password@tcp(localhost:3307)/twitter_clone?charset=utf8mb4&parseTime=True&loc=Local")
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
