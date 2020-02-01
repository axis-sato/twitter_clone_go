package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mattn/go-gimei"
	"github.com/volatiletech/sqlboiler/queries"
)

const userLength = 20

func makeDummyUsers(ctx context.Context, db *sql.DB) {
	v := ""
	for i := 0; i < userLength; i++ {
		name := gimei.NewName().Kanji()
		icon := "https://avatars2.githubusercontent.com/u/1905224?s=460&v=4"
		profile := fmt.Sprintf("こんにちは。%sと申します。\nよろしくお願いします。", name)

		v += fmt.Sprintf("('%s','%s','%s',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)", name, icon, profile)

		if i == userLength-1 {
			v += ";"
		} else {
			v += ","
		}
	}

	if _, err := queries.Raw("INSERT INTO users (name, icon, profile, created_at, updated_at) VALUES "+v).ExecContext(ctx, db); err != nil {
		panic(err.Error())
	}
}
