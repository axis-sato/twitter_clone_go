package main

//go:generate sqlboiler --wipe --no-tests mysql

import (
	"fmt"

	"github.com/c8112002/twitter_clone_go/router"

	"github.com/c8112002/twitter_clone_go/db"
	"github.com/c8112002/twitter_clone_go/handler"
	"github.com/c8112002/twitter_clone_go/store"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	d, err := db.New(true)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if err := d.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	e := router.New()

	us := store.NewUserStore(d)

	h := handler.NewHandler(us)

	v1 := e.Group("/api/v1")
	h.Register(v1)

	e.Logger.Fatal(e.Start(":1323"))
}
