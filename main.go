package main

//go:generate sqlboiler --wipe --no-tests mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/c8112002/twitter_clone_go/router"

	"github.com/c8112002/twitter_clone_go/db"
	"github.com/c8112002/twitter_clone_go/handler"
	"github.com/c8112002/twitter_clone_go/store"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loc, err := time.LoadLocation("Asia/Tokyo")

	time.Local = loc

	if err != nil {
		fmt.Println(err.Error())
	}

	d, err := db.New(true, loc)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if err := d.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	e := router.New()

	ctx := context.Background()
	us := store.NewUserStore(d, ctx)
	ts := store.NewTweetStore(d, ctx)

	h := handler.NewHandler(us, ts)

	h.Register(e)

	e.Logger.Fatal(e.Start(":1323"))
}
