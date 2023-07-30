package main

import (
	"morph_mails/ent"
	"morph_mails/ent/connector"
	"morph_mails/ent/receiver"
	"morph_mails/ent/user"
	"net/http"

	"github.com/google/uuid"
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

func connectAccount(c framework.Context) error {
	id := c.Param(fieldID)
	uniqueID := c.QueryParam("uid")

	if len(uniqueID) == 0 {
		return StatusReport(c, http.StatusBadRequest)
	}

	connectorID := c.Get("connectorID").(string)
	parsedUUID, err := uuid.Parse(uniqueID)
	if err != nil {
		return StatusReport(c, http.StatusBadRequest)
	}

	_, err = db.Receiver.Create().SetUserID(parsedUUID).SetConnectorID(connectorID).SetID(id).Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "connectAccount",
		}).Error(err)

		return StatusReport(c, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"code":    http.StatusOK,
		"message": "Successfully connected a new Receiver(connector_id = " + connectorID + ") to User(id = " + id + ")",
	})
}

func getUniqueID(c framework.Context) error {
	id := c.Param("id")

	userLoaded, err := db.User.Query().Where(user.HasReceiversWith(receiver.ID(id))).Select("id").First(ctx)
	if ent.IsNotFound(err) {
		return StatusReport(c, http.StatusNotFound)
	} else if err != nil {
		log.WithFields(log.Fields{
			"function": "getUniqueID",
		}).Error(err)
		return StatusReport(c, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"id": userLoaded.ID,
	})
}

func isRegistered(c framework.Context) error {
	id := c.Param("id")

	exist, err := db.User.Query().Where(user.HasReceiversWith(receiver.ID(id))).Exist(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "isRegistered",
		}).Error(err)

		return StatusReport(c, http.StatusInternalServerError)
	}

	if !exist {
		return c.JSON(http.StatusNotFound, map[string]any{
			"code": http.StatusNotFound,
			"message": "User with that ID wasn't found among any of connectors. " +
				"Perhaps you need to connect yourself to existing user or create a new one with routes below",
			"routes": map[string]any{
				"create":        "/v2/register/" + id,
				"get_unique_id": "/v2/unique_id/" + id,
				"connect":       "/v2/connect/<unique_id>",
			},
		})
	}

	return c.String(http.StatusOK, "OK")
}

func registerAccount(c framework.Context) error {
	id := c.Param(fieldID)
	connectorID := c.Get(fieldConID).(string)

	exist, err := db.User.Query().Where(user.HasReceiversWith(receiver.ID(id))).Exist(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "isRegistered",
		}).Error(err)

		return StatusReport(c, http.StatusInternalServerError)
	}

	if exist {
		return StatusReport(c, http.StatusFound)
	}

	userCreated, err := db.User.Create().AddReceivers().Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "isRegistered",
		}).Error(err)

		return StatusReport(c, http.StatusInternalServerError)
	}

	err = db.Receiver.Create().SetID(id).SetUser(userCreated).SetConnectorID(connectorID).Exec(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "isRegistered",
		}).Error(err)

		return StatusReport(c, http.StatusInternalServerError)
	}

	log.WithFields(log.Fields{
		"userID":      id,
		"connectorID": connectorID,
	}).Infoln("New user registered in the bot")

	return c.JSON(http.StatusOK, userCreated)
}

func assignNew(c framework.Context) error {
	id := c.Param(fieldID)

	genID, err := sid.Generate()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "assignNew(sid.Generate)",
		}).Error(err)

		return StatusReport(c, http.StatusInternalServerError)
	}

	generated := genID + "@" + osDomain

	_, err = db.User.Update().Where(user.HasReceiversWith(receiver.ID(id))).AppendEmails([]string{generated}).Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorln("Caught an error while pushing assigned email to a DB")

		return StatusReport(c, http.StatusBadRequest)
	}

	log.WithFields(log.Fields{
		"generatedEmail": generated,
		"ID":             c.Param("id"),
	}).Infoln("Successfully assigned new email")

	return c.JSON(http.StatusOK, map[string]any{
		"id":    c.Param("id"),
		"email": generated,
	})
}

func delAll(c framework.Context) error {
	id := c.Param(fieldID)

	err := db.User.Update().Where(user.HasReceiversWith(receiver.ID(id))).SetEmails([]string{}).Exec(ctx)

	if ent.IsNotFound(err) {
		return StatusReport(c, http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": c.Param("id"),
	}).Infoln("Cleared all email addresses")

	return StatusReport(c, http.StatusOK)
}

func delSome(c framework.Context) error {
	id := c.Param(fieldID)
	email := c.Param("email")

	userLoaded, err := db.User.Query().Where(user.HasReceiversWith(receiver.ID(id))).First(ctx)
	if ent.IsNotFound(err) || len(userLoaded.Emails) == 0 || !slices.Contains(userLoaded.Emails, email) {
		return StatusReport(c, http.StatusNotFound)
	}

	err = db.User.Update().
		Where(user.HasReceiversWith(receiver.ID(id))).
		SetEmails(remove(userLoaded.Emails, email)).
		Exec(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function":   "delSome",
			"email":      email,
			"userLoaded": userLoaded,
			"error":      err,
		}).Errorln("Failed to delete an email address")

		return StatusReport(c, http.StatusInternalServerError)
	}

	log.WithFields(log.Fields{
		"id":    c.Param("id"),
		"email": email,
	}).Infoln("Deleted an email address")

	return StatusReport(c, http.StatusOK)
}

func listAll(c framework.Context) error {
	id := c.Param(fieldID)
	connectorID := c.Get(fieldConID).(string)

	resp, err := db.User.Query().
		WithReceivers().
		Where(
			user.HasReceiversWith(
				receiver.And(
					receiver.ID(id),
					receiver.HasConnectorWith(connector.ID(connectorID)),
				),
			),
		).
		First(ctx)

	if ent.IsNotFound(err) || len(resp.Emails) == 0 {
		return StatusReport(c, http.StatusNotFound)
	} else if err != nil {
		log.WithFields(log.Fields{
			"function": "listAll",
		}).Error(err)
		return StatusReport(c, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"emails": resp.Emails,
	})
}

func turnFwd(c framework.Context) error {
	id := c.Param(fieldID)
	connectorID := c.Get(fieldConID).(string)

	getUser, err := db.User.Query().
		Where(
			user.HasReceiversWith(
				receiver.And(
					receiver.ID(id),
					receiver.HasConnectorWith(connector.ID(connectorID)),
				),
			),
		).
		First(ctx)
	if err != nil {
		return StatusReport(c, http.StatusInternalServerError)
	}

	err = getUser.Update().SetForward(!getUser.Forward).Exec(ctx)
	if err != nil {
		return StatusReport(c, http.StatusInternalServerError)
	}

	log.WithFields(log.Fields{
		"id":      id,
		"forward": !getUser.Forward,
	}).Infoln("Switched forward mode")

	return c.JSON(http.StatusOK, map[string]any{
		"id":      id,
		"forward": !getUser.Forward,
	})
}

func sendAnnouncement(c framework.Context) error {
	subject := c.FormValue("Subject")
	text := c.FormValue("text")
	url := c.FormValue("url")

	switch {
	case len(subject) == 0:
		subject = "[No Subject]"
	case len(text) == 0:
		text = "[No Body]"
	case len(url) == 0:
		url = "https://www." + osDomain
	}

	users, _ := db.User.Query().Where(user.Forward(true)).All(ctx)

	for _, u := range users {
		for _, r := range u.Edges.Receivers {
			finalRes := result{
				Message: message{
					From: "OfficialNewsletter@" + osDomain,
					To:   u.Emails[0],
					Text: text,
				},
				RenderedURI: url,
				ID:          r.ID,
			}

			_, err := req.R().
				SetHeader("Content-type", "application/json").
				SetBodyJsonMarshal(finalRes).
				SetHeader("user-agent", "github.com/voxelin").
				Post(r.Edges.Connector.URL)
			if err != nil {
				return StatusReport(c, http.StatusBadRequest)
			}
		}
	}

	log.WithFields(log.Fields{
		"Subject": subject,
		"text":    text,
		"url":     url,
	}).Infoln("Sent an announcement")

	return c.String(http.StatusOK, "OK")
}
