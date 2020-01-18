package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/volatiletech/sqlboiler/queries"
)

func makeDummyLikes(ctx context.Context, db *sql.DB)  {
	v := ""
	l := userLength / 2
	for i := 0; i < l; i++ {
		tis, err := faker.RandomInt(1, tweetLength - 1)

		if err != nil {
			panic(err.Error())
		}

		ui := i + 1

		for j, ti := range tis[:tweetLength / 2] {
			v += fmt.Sprintf("('%d','%d',CURRENT_TIMESTAMP)", ti, ui)

			if i == l - 1 && j == tweetLength / 2 - 1 {
				v += ";"
			} else {
				v += ","
			}
		}
	}

	if _, err := queries.Raw("INSERT INTO likes (tweet_id, user_id, created_at) VALUES " + v).ExecContext(ctx, db); err != nil {
		panic(err.Error())
	}
}