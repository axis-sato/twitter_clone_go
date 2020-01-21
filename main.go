package main

//go:generate sqlboiler --wipe --no-tests mysql

import (
	"fmt"
	db2 "github.com/c8112002/twitter_clone_go/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {

	db, err := db2.New()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	e := echo.New()
	e.GET("/", users)
	e.Logger.Fatal(e.Start(":1234"))
}

func users(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
