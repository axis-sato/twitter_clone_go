package db

import (
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/spf13/viper"
)

func New(debug bool, loc *time.Location) (*sql.DB, error) {
	c, err := readDBConf()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open(c.Development.Dialect, c.Development.Datasource)

	if err != nil {
		return db, err
	}

	boil.DebugMode = debug
	boil.SetLocation(loc)

	return db, nil
}

func TestDB(debug bool, loc *time.Location) (*sql.DB, error) {
	c, err := readDBConf()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open(c.Test.Dialect, c.Test.Datasource)

	if err != nil {
		return db, err
	}

	boil.DebugMode = debug
	boil.SetLocation(loc)

	return db, nil
}

func readDBConf() (*dbconf, error) {
	var c dbconf

	viper.SetConfigName("dbconf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("./db")
	viper.AddConfigPath("../db")

	if err := viper.ReadInConfig(); err != nil {
		return &c, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return &c, err
	}

	return &c, nil
}

type dbconf struct {
	Development param
	Test        param
}

type param struct {
	Dialect    string
	Datasource string
	Dir        string
}
