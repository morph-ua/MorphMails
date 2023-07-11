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

func assignNew(c framework.Context) error {
	id := c.Param("id")

	var result Account
	db.Where(&Account{ID: id}).First(&result)

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

	res := db.Model(&Account{}).Where(&Account{ID: id}).Update("emails", gorm.Expr("emails || ?", pq.StringArray{generated}))

	if res.Error != nil {
		log.WithFields(log.Fields{
			"error": res.Error.Error(),
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

	response := db.Where(&Account{ID: id})

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &HttpError{
			Status:  http.StatusNotFound,
			Message: "User with ID: `" + id + "` wasn't found in API's database. Please, register yourself in the bot",
		})
	}

	db.Model(&Account{}).Where(&Account{ID: id}).Update("emails", pq.StringArray{id + "@decline.live"})

	log.WithFields(log.Fields{
		"id": id,
	}).Infoln("Cleared all email addresses")

	return c.String(http.StatusOK, "OK")
}

func delSome(c framework.Context) error {
	id := c.Param("id")

	var result Account
	response := db.Where(&Account{ID: id}).First(&result)

	email := c.Param("email")

	if len(email) == 0 || email == result.ID+"@decline.live" {
		return c.JSON(http.StatusBadRequest, &HttpError{
			Status:  http.StatusBadRequest,
			Message: "You can not delete your persistent email address",
		})
	}

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, notFoundMessage)
	}

	res := db.Model(&Account{}).Where(&Account{ID: id}).Update("emails", gorm.Expr("ARRAY_REMOVE(emails, ?)", email))

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
	response := db.Where(&Account{ID: id}).Find(&result)

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, notFoundMessage)
	}

	return c.JSON(http.StatusOK, result)
}

func turnFwd(c framework.Context) error {
	id := c.Param("id")

	var result Account
	tx := db.Model(&Account{}).Where(&Account{ID: id}).Update("forward", gorm.Expr("NOT forward")).Select("forward").First(&result)

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

func registerNew(c framework.Context) error {
	id := c.Param("id")
	clientID := c.Get("client").(string)

	var result Account
	resp := db.Where("id = ?", id).Select("clients").First(&result)

	if errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		created := &Account{
			ID:      id,
			Emails:  pq.StringArray{id + "@decline.live"},
			Clients: pq.StringArray{clientID},
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

	if !slices.Contains(result.Clients, clientID) {
		db.Model(&Account{}).Where("id = ?", id).Update("clients", gorm.Expr("clients || ?", pq.StringArray{clientID}))

		log.WithFields(log.Fields{
			"userID": id,
			"client": clientID,
		}).Infoln("Client connected to an account")

		return c.String(http.StatusOK, "OK")
	}

	return c.String(http.StatusFound, "Found")
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

	var result []Account
	db.Select("id", "clients").Find(&result)

	for _, account := range result {
		finalRes := FinalResult{
			Message:     fmt.Sprintf(paidMessageTemplate, "news@atomic.decline.live", account.ID+"@decline.live", subject, text),
			RenderedURI: url,
			ID:          account.ID,
		}

		for _, clientID := range account.Clients {
			client := getClient(clientID)
			if len(client) == 0 {
				return c.String(http.StatusNotFound, "Client wasn't found in DB")
			}

			_, err := reqClient.R().
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
