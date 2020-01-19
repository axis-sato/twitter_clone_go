package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {

	c := readDBConf()

	db, err := sql.Open(c.Development.Driver, c.Development.Open)
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

func readDBConf() *dbconf {
	var c dbconf

	viper.SetConfigName("dbconf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("./db")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err.Error())
	}

	return &c
}

type dbconf struct {
	Development struct{
		Driver string
		Open string
	}
}
