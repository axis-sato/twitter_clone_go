package main

import (
	"context"
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/mattn/go-gimei"
	"github.com/volatiletech/sqlboiler/queries"
)

func makeDummyUsers(ctx context.Context, db *sql.DB)  {
	v := ""
	l := 20
	for i := 0; i < l; i++ {
		name := gimei.NewName().Kanji()
		icon := faker.URL()
		profile := "こんにちは。" + name + "と申します。\nよろしくお願いします。"

		v += "('" + name + "','" + icon + "','" + profile + "',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)"

		if i == l - 1 {
			v += ";"
		} else {
			v += ","
		}
	}

	if _, err := queries.Raw("INSERT INTO users (name, icon, profile, created_at, updated_at) VALUES " + v).ExecContext(ctx, db); err != nil {
		panic(err.Error())
	}
}