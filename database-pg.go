//go:build !database_oci8 && !database_mysql
// +build !database_oci8,!database_mysql

package main

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func setupDB(s string) (*sql.DB, error) {
	return sql.Open("pq", s)
}

func setupBunDB(sqldb *sql.DB) (*bun.DB, error) {
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db, nil
}
