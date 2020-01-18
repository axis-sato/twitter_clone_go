package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/volatiletech/sqlboiler/queries"
)

func makeDummyFollowers(ctx context.Context, db *sql.DB)  {
	v := ""
	l := userLength / 2
	for i := 0; i < l; i++ {
		is, err := faker.RandomInt(1, userLength - 1)

		if err != nil {
			panic(err.Error())
		}

		followerId := i + 1

		fl := 10
		j := 0
		for _, fi := range is {

			if fi == followerId {
				continue
			}

			v += fmt.Sprintf("('%d','%d',CURRENT_TIMESTAMP)", followerId, fi)

			if i == l - 1 && j == fl - 1 {
				v += ";"
			} else {
				v += ","
			}

			if j == fl - 1 {
				break
			}

			j++
		}
	}

	if _, err := queries.Raw("INSERT INTO followers (follower_id, followee_id, created_at) VALUES " + v).ExecContext(ctx, db); err != nil {
		panic(err.Error())
	}
}