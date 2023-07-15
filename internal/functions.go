package main

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/rand"
	"net/http"
)

//goland:noinspection ALL
func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(osDB), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Silent),
		PrepareStmt: true,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"fatal":    true,
			"function": "connect",
			"error":    err,
		}).Fatalln("Database connection aborted")
	}

	tx := db.Session(&gorm.Session{PrepareStmt: true})

	return tx
}

func syncWithClients(json FinalResult, id string) int {
	client := getClient(id)
	if len(client) == 0 {
		return http.StatusBadRequest
	}

	r, err := req.R().
		SetHeader("Content-type", "application/json").
		SetBodyJsonMarshal(json).
		SetHeader("user-agent", "github.com/voxelin").
		Post(client)

	if err != nil {
		return http.StatusInternalServerError
	}

	return r.StatusCode
}
