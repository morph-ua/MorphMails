package main

import (
	"database/sql"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"helium/ent"
)

func StatusReport(ctx echo.Context, c int) error {
	return ctx.JSON(c, map[string]any{
		"code":    c,
		"message": http.StatusText(c),
	})
}

// Why doesn't Golang have a function to delete an element from a slice???
func remove[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}

	return l
}

// Open database connection with URL and not DSN
func Open(url string) *ent.Client {
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "init",
		}).Fatalln("Failed to open database connection.")
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	return ent.NewClient(ent.Driver(drv))
}

// This function sends requests and does nothing more.
// TODO: Add status code handler to notify of failure
func syncConnectors(json result, receivers []*ent.Receiver) {
	for _, receiver := range receivers {
		json.ID = receiver.ID
		_, _ = req.R().
			SetHeader("Content-type", "application/json").
			SetBodyJsonMarshal(json).
			SetHeader("user-agent", "github.com/voxelin").
			Post(receiver.Edges.Connector.URL)
	}
}
