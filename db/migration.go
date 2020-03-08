package db

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateTestDB(db *sql.DB) error {
	return exec(db, migrate.Up)
}

func DropTestDB(db *sql.DB) error {
	return exec(db, migrate.Down)
}

func exec(db *sql.DB, dir migrate.MigrationDirection) error {
	m := &migrate.FileMigrationSource{Dir: "../db/migrations"}
	_, err := migrate.Exec(db, "mysql", m, dir)
	return err
}
