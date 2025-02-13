//go:build database_oci8

package main

import (
	"database/sql"

	_ "github.com/mattn/go-oci8"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/oracledialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func setupDB(s string) (*sql.DB, error) {
	return sql.Open("oci8", s)
}

func setupBunDB(sqldb *sql.DB) (*bun.DB, error) {
	db := bun.NewDB(sqldb, oracledialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db, nil
}
