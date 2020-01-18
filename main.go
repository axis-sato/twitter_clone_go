package main

//go:generate sqlboiler --wipe --no-tests mysql

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
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

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	e := echo.New()
	e.GET("/", users)
	e.Logger.Fatal(e.Start(":1234"))
}

func users(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
