package main

import (
	"errors"
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func uploadHTML(html string, from string, to string) string {
	if len(html) == 0 || len(from) == 0 || len(to) == 0 {
		log.WithFields(log.Fields{"html": html, "from": from, "to": to}).Errorln("Did not upload letter due to lack of values")
		return ""
	}

	id := randSeq(8)
	for {
		var count int64
		db.Model(&Letter{}).Where("id = ?", id).Count(&count)
		if count == 0 {
			break
		}
		id = randSeq(8)
	}
	data := Letter{
		ID:   id,
		Html: html,
		From: from,
		To:   to,
	}
	db.Create(&data)

	return id
}

func getRaw(c framework.Context) error {
	id := c.Param("id")

	var result Letter
	res := db.Where("id = ?", id).First(&result)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, Letter{
			ID:   id,
			Html: "<html><head><title>404 Not Found</title></head><body><center><h1>404 Not Found</h1><hr>Atomic Emails</center></body>",
			From: "not.found@db",
			To:   "not.found@db",
		})
	}

	return c.JSON(http.StatusOK, result)
}

func getHTML(c framework.Context) error {
	id := c.Param("id")

	var result Letter
	res := db.Where("id = ?", id).First(&result)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.HTML(http.StatusNotFound, "<html><head><title>404 Not Found</title></head><body><center><h1>404 Not Found</h1><hr>Atomic Emails</center></body>")
	}

	return c.HTML(http.StatusOK, result.Html)
}
