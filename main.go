package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64      `bun:"id,pk,autoincrement" form:"id"`
	Content   string     `bun:"content,notnull" form:"content"`
	Done      bool       `bun:"done" form:"done"`
	Until     *time.Time `bun:"until" from:"until"`
	DeletedAt *time.Time `bun:",soft_delete"`
}

func main() {
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	ctx := context.Background()
	_, err = db.NewCreateTable().Model((*Todo)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
