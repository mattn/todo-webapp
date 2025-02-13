//go:build database_mysql

package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func setupDB(s string) (*sql.DB, error) {
	return sql.Open("mysql", s)
}

func setupBunDB(sqldb *sql.DB) (*bun.DB, error) {
	db := bun.NewDB(sqldb, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db, nil
}
