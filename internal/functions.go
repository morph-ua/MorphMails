package main

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"helium/ent"
	"math/rand"
	"net/http"
)

func StatusReport(ctx echo.Context, c int) error {
	return ctx.JSON(c, map[string]interface{}{
		"code":    c,
		"message": http.StatusText(c),
	})
}

// Generate a random sequence of characters (Currently used to generate email accounts, may be replaced later)
//
//goland:noinspection ALL
func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
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
func syncConnectors(json Result, url string) {
	_, _ = req.R().
		SetHeader("Content-type", "application/json").
		SetBodyJsonMarshal(json).
		SetHeader("user-agent", "github.com/voxelin").
		Post(url)
}
