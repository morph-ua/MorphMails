package main

import (
	"context"
	"errors"
	"fmt"
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"helium/ent/receiver"
	"helium/ent/user"
	"net/http"
	"strconv"
)

//func registerNew(c framework.Context) error {
//	id := c.Param("id")
//	client := c.Get("client").(string)
//
//	resp, err := db.User.Query().WithReceivers().Where(user.HasReceiversWith(receiver.ID(id))).First(context.Background())
//
//	if ent.IsNotFound(err) {
//		created := &Account{
//			Receivers: datatypes.JSONSlice[Receiver]{{
//				ID:     id,
//				Client: client,
//			}},
//		}
//
//		res := db.Create(created)
//
//		if res.Error != nil {
//			log.Error(res.Error)
//			return c.JSON(http.StatusInternalServerError, internalServerErrorMessage)
//		}
//
//		log.WithFields(log.Fields{
//			"userID": id,
//			"client": client,
//		}).Infoln("New user registered in the bot")
//
//		return c.JSON(http.StatusOK, created)
//	}
//
//	if !slices.Contains(account.Receivers, Receiver{ID: id, Client: client}) {
//		account.Receivers = append(account.Receivers, Receiver{ID: id, Client: client})
//		db.Model(&account).Updates(Account{Receivers: account.Receivers})
//
//		log.WithFields(log.Fields{
//			"userID": id,
//			"client": client,
//		}).Infoln("Client connected to an account")
//
//		return c.String(http.StatusOK, "OK")
//	}
//
//	return c.String(http.StatusFound, "Found")
//}

func assignNew(c framework.Context) error {
	id := c.Param("id")

	account, err := db.User.Query().WithReceivers().Where(user.HasReceiversWith(receiver.ID(id))).First(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error. - "+err.Error())
	}

	_, err = strconv.Atoi(id)
	if len(id) == 0 || err != nil {
		return c.JSON(http.StatusBadRequest, badRequestMessage)
	}

	generated := randSeq(8) + "@decline.live"

loop:
	for _, email := range account.Emails {
		if email == generated {
			generated = randSeq(8) + "@decline.live"
			goto loop
		}
	}

	_, err = db.User.Update().Where(user.HasReceiversWith(receiver.ID(id))).AppendEmails([]string{generated}).Save(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorln("Caught an error while pushing assigned email to a DB")

		return c.JSON(http.StatusBadRequest, badRequestMessage)
	}

	log.WithFields(log.Fields{
		"generatedEmail": generated,
		"ID":             c.Param("id"),
	}).Infoln("Successfully assigned new email")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":    c.Param("id"),
		"email": generated,
	})
}

func delAll(c framework.Context) error {
	id := c.Param("id")

	response := db.Where("receivers @> ?", datatypes.JSONSlice[map[string]interface{}]{{"id": id}})

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &HttpError{
			Status:  http.StatusNotFound,
			Message: "User with ID: `" + c.Param("id") + "` wasn't found in API's database. Please, register yourself in the bot",
		})
	}

	db.Model(&Account{}).Where("receiver @> ?", datatypes.JSONSlice[map[string]interface{}]{{"id": id}}).Update("emails", nil)

	log.WithFields(log.Fields{
		"id": c.Param("id"),
	}).Infoln("Cleared all email addresses")

	return c.String(http.StatusOK, "OK")
}

func delSome(c framework.Context) error {
	id := c.Param("id")

	var result Account
	response := db.Where("receivers @> ?", datatypes.JSONSlice[map[string]interface{}]{{"id": id}}).Find(&result)

	email := c.Param("email")

	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, notFoundMessage)
	}

	res := db.Model(&Account{}).Where("receivers @> ?", datatypes.JSONSlice[map[string]interface{}]{{"id": id}}).Update("emails", gorm.Expr("ARRAY_REMOVE(emails, ?)", email))

	if res.Error != nil {
		log.Error(res.Error)
		return c.JSON(http.StatusInternalServerError, internalServerErrorMessage)
	}

	log.WithFields(log.Fields{
		"id":    c.Param("id"),
		"email": email,
	}).Infoln("Deleted an email address")

	return c.String(http.StatusOK, "OK")
}

func listAll(c framework.Context) error {
	id := c.Param("id")
	client := c.Get("client").(string)

	var a []Account
	response := db.Debug().Joins("JOIN receivers ON receivers.id = ? AND receivers.client = ?", id, client).Find(&a)
	println(a)
	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, notFoundMessage)
	}

	return c.JSON(http.StatusOK, a)
}

func turnFwd(c framework.Context) error {
	id := c.Param("id")

	var result Account
	tx := db.Model(&Account{}).Where("receivers @> ?", datatypes.JSONSlice[map[string]interface{}]{{"id": id}}).Update("forward", gorm.Expr("NOT forward")).Select("forward").First(&result)

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
	db.Select("receivers").Find(&accounts)

	for _, account := range accounts {
		for _, receiver := range account.Receivers {
			finalRes := FinalResult{
				Message:     fmt.Sprintf(messageTemplate, "hello@lowt.live", "news_receiver@decline.live", subject, text),
				RenderedURI: url,
				ID:          receiver.ID,
			}

			client := getConnector(receiver.Client)
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
