package db

import (
	"database/sql"
	"github.com/spf13/viper"
)

func New() (*sql.DB, error) {
	c, err := readDBConf()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open(c.Development.Driver, c.Development.Open)

	if err != nil {
		return db, err
	}

	return db, nil
}

func readDBConf() (*dbconf, error) {
	var c dbconf

	viper.SetConfigName("dbconf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("./db")

	if err := viper.ReadInConfig(); err != nil {
		return &c, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return &c, err
	}

	return &c, nil
}

type dbconf struct {
	Development struct{
		Driver string
		Open string
	}
}
