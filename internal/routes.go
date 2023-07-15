package main

import (
	"errors"
	"fmt"
	framework "github.com/labstack/echo/v4"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func registerNew(c framework.Context) error {
	id := c.Param("id")
	clientID := c.Get("client").(string)

	var account Account
	a := `[ { "id": ` + strconv.Quote(id) + ` } ]`
	resp := db.Where("destination @> ?", a).Find(&account)

	if errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		created := &Account{
			Destination: []Destination{{
				ID:     id,
				Client: clientID,
			}},
		}

		res := db.Create(created)

		if res.Error != nil {
			log.Error(res.Error)
			return c.JSON(http.StatusInternalServerError, internalServerErrorMessage)
		}

		log.WithFields(log.Fields{
			"userID": id,
			"client": clientID,
		}).Infoln("New user registered in the bot")

		return c.JSON(http.StatusOK, created)
	}

	if !slices.Contains(account.Destination, Destination{ID: id, Client: clientID}) {
		account.Destination = append(account.Destination, Destination{ID: id, Client: clientID})
		db.Model(&account).Updates(Account{Destination: account.Destination})

		log.WithFields(log.Fields{
			"userID": id,
			"client": clientID,
		}).Infoln("Client connected to an account")

		return c.String(http.StatusOK, "OK")
	}

	return c.String(http.StatusFound, "Found")
}

func assignNew(c framework.Context) error {
	id := c.Param("id")

	var result Account
	a := `[ { "id": ` + strconv.Quote(id) + ` } ]`
	db.Where("destination @> ?", a).Find(&result).First(&result)

	_, err := strconv.Atoi(id)
	if len(id) == 0 || err != nil {
		return c.JSON(http.StatusBadRequest, badRequestMessage)
	}

	generated := randSeq(8) + "@decline.live"

loop:
	for _, email := range result.Emails {
		if email == generated {
			generated = randSeq(8) + "@decline.live"
			goto loop
		}
	}

	res := db.Model(&Account{}).Where("destination @> ?", a).Update("emails", gorm.Expr("emails || ?", pq.StringArray{generated}))

	if res.Error != nil {
		log.WithFields(log.Fields{
			"error": res.Error,
		}).Errorln("Caught an error while pushing assigned email to a DB")

		return c.JSON(http.StatusBadRequest, badRequestMessage)
	}

	log.WithFields(log.Fields{
		"generatedEmail": generated,
		"ID":             id,
	}).Infoln("Successfully assigned new email")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"email": generated,
	})
}

func delAll(c framework.Context) error {
	id := c.Param("id")

	a := `[ { "id": ` + strconv.Quote(id) + ` } ]`
	response := db.Where("destination @> ?", a)

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &HttpError{
			Status:  http.StatusNotFound,
			Message: "User with ID: `" + id + "` wasn't found in API's database. Please, register yourself in the bot",
		})
	}

	db.Model(&Account{}).Where("destination @> ?", a).Update("emails", nil)

	log.WithFields(log.Fields{
		"id": id,
	}).Infoln("Cleared all email addresses")

	return c.String(http.StatusOK, "OK")
}

func delSome(c framework.Context) error {
	id := c.Param("id")

	var result Account
	a := `[ { "id": ` + strconv.Quote(id) + ` } ]`
	response := db.Where("destination @> ?", a).Find(&result)

	email := c.Param("email")

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, notFoundMessage)
	}

	res := db.Model(&Account{}).Where("destination @> ?", a).Update("emails", gorm.Expr("ARRAY_REMOVE(emails, ?)", email))

	if res.Error != nil {
		log.Error(res.Error)
		return c.JSON(http.StatusInternalServerError, internalServerErrorMessage)
	}

	log.WithFields(log.Fields{
		"id":    id,
		"email": email,
	}).Infoln("Deleted an email address")

	return c.String(http.StatusOK, "OK")
}

func listAll(c framework.Context) error {
	id := c.Param("id")

	var result Account
	a := `[ { "id": ` + strconv.Quote(id) + ` } ]`
	response := db.Where("destination @> ?", a).Select("emails").First(&result)

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, notFoundMessage)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"emails": result.Emails})
}

func turnFwd(c framework.Context) error {
	id := c.Param("id")

	var result Account
	a := `[ { "id": ` + strconv.Quote(id) + ` } ]`
	tx := db.Model(&Account{}).Where("destination @> ?", a).Update("forward", gorm.Expr("NOT forward")).Select("forward").First(&result)

	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, &HttpError{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	log.WithFields(log.Fields{
		"id":      id,
		"forward": result.Forward,
	}).Infoln("Switched forward mode")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"forward": result.Forward,
	})
}

func sendAnnouncement(c framework.Context) error {
	subject := c.FormValue("subject")
	text := c.FormValue("text")
	url := c.FormValue("url")

	switch {
	case len(subject) == 0:
		subject = "[No Subject]"
	case len(text) == 0:
		text = "[No Body]"
	case len(url) == 0:
		url = "https://www.decline.live"
	}

	var accounts []Account
	db.Select("destination").Find(&accounts)

	for _, account := range accounts {
		for _, destination := range account.Destination {
			finalRes := FinalResult{
				Message:     fmt.Sprintf(messageTemplate, "hello@lowt.live", "news_receiver@decline.live", subject, text),
				RenderedURI: url,
				ID:          destination.ID,
			}

			client := getClient(destination.Client)
			if len(client) == 0 {
				return c.String(http.StatusNotFound, "Client wasn't found in DB")
			}

			_, err := req.R().
				SetHeader("Content-type", "application/json").
				SetBodyJsonMarshal(finalRes).
				SetHeader("user-agent", "github.com/voxelin").
				Post(client)

			if err != nil {
				return c.JSON(http.StatusBadRequest, badRequestMessage)
			}
		}
	}

	log.WithFields(log.Fields{
		"subject": subject,
		"text":    text,
		"url":     url,
	}).Infoln("Sent an announcement")

	return c.String(http.StatusOK, "OK")
}
