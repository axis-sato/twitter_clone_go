package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/volatiletech/sqlboiler/queries"
)

const tweetLength = 100

func makeDummyTweets(ctx context.Context, db *sql.DB) {
	v := ""
	for i := 0; i < tweetLength; i++ {
		is, err := faker.RandomInt(1, userLength-1)

		if err != nil {
			panic(err.Error())
		}

		id := i + 1
		userID := is[0]

		tweet := fmt.Sprintf("適当なツイート(id = %d)", id)

		v += fmt.Sprintf("('%d', '%s',CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)", userID, tweet)

		if i == tweetLength-1 {
			v += ";"
		} else {
			v += ","
		}
	}

	if _, err := queries.Raw("INSERT INTO tweets (user_id, tweet, created_at, updated_at) VALUES "+v).ExecContext(ctx, db); err != nil {
		panic(err.Error())
	}
}
