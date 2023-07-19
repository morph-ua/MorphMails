package main

import (
	"github.com/google/uuid"
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"helium/ent"
	"helium/ent/connector"
	"helium/ent/receiver"
	"helium/ent/user"
	"net/http"
	"strconv"
)

func connectAccount(c framework.Context) error {
	id := c.Param("id")
	uniqueID := c.QueryParam("uid")
	if len(uniqueID) == 0 {
		return StatusReport(c, 400)
	}

	connectorID := c.Get("connectorID").(string)
	parsedUUID, err := uuid.Parse(uniqueID)
	if err != nil {
		return StatusReport(c, 400)
	}

	_, err = db.Receiver.Create().SetUserID(parsedUUID).SetConnectorID(connectorID).SetID(id).Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "connectAccount",
		}).Error(err)
		return StatusReport(c, 500)
	}

	return c.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "Successfully connected a new Receiver(connector_id = " + connectorID + ") to User(id = " + id + ")",
	})
}

func getUniqueID(c framework.Context) error {
	id := c.Param("id")

	userLoaded, err := db.User.Query().Where(user.HasReceiversWith(receiver.ID(id))).Select("id").First(ctx)
	if ent.IsNotFound(err) {
		return StatusReport(c, 404)
	} else if err != nil {
		log.WithFields(log.Fields{
			"function": "getUniqueID",
		}).Error(err)
		return StatusReport(c, 500)
	}

	return c.JSON(200, map[string]interface{}{
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
		return StatusReport(c, 500)
	}

	if !exist {
		return c.JSON(404, map[string]interface{}{
			"code":    404,
			"message": "User with that ID wasn't found among any of connectors. Perhaps you need to connect yourself to existing user or create a new one with routes below",
			"routes": map[string]interface{}{
				"create":        "/v2/register/" + id,
				"get_unique_id": "/v2/unique_id/" + id,
				"connect":       "/v2/connect/<unique_id>",
			},
		})
	}

	return c.String(200, "OK")
}

func registerAccount(c framework.Context) error {
	id := c.Param("id")
	connectorID := c.Get("connectorID").(string)

	exist, err := db.User.Query().Where(user.HasReceiversWith(receiver.ID(id))).Exist(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "isRegistered",
		}).Error(err)
		return StatusReport(c, 500)
	}

	if exist {
		return StatusReport(c, 302)
	}

	userCreated, err := db.User.Create().AddReceivers().Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "isRegistered",
		}).Error(err)
		return StatusReport(c, 500)
	}

	err = db.Receiver.Create().SetID(id).SetUser(userCreated).SetConnectorID(connectorID).Exec(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "isRegistered",
		}).Error(err)
		return StatusReport(c, 500)
	}

	log.WithFields(log.Fields{
		"userID":      id,
		"connectorID": connectorID,
	}).Infoln("New user registered in the bot")

	return c.JSON(http.StatusOK, userCreated)
}

func assignNew(c framework.Context) error {
	id := c.Param("id")

	account, err := db.User.Query().WithReceivers().Where(user.HasReceiversWith(receiver.ID(id))).First(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error. - "+err.Error())
	}

	_, err = strconv.Atoi(id)
	if len(id) == 0 || err != nil {
		return StatusReport(c, 400)
	}

	generated := randSeq(8) + "@" + osDomain

loop:
	for _, email := range account.Emails {
		if email == generated {
			generated = randSeq(8) + "@" + osDomain
			goto loop
		}
	}

	_, err = db.User.Update().Where(user.HasReceiversWith(receiver.ID(id))).AppendEmails([]string{generated}).Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorln("Caught an error while pushing assigned email to a DB")

		return StatusReport(c, 400)
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

	err := db.User.Update().Where(user.HasReceiversWith(receiver.ID(id))).SetEmails([]string{}).Exec(ctx)

	if ent.IsNotFound(err) {
		return StatusReport(c, 404)
	}

	log.WithFields(log.Fields{
		"id": c.Param("id"),
	}).Infoln("Cleared all email addresses")

	return StatusReport(c, 200)
}

func delSome(c framework.Context) error {
	id := c.Param("id")
	email := c.Param("email")

	userLoaded, err := db.User.Query().Where(user.HasReceiversWith(receiver.ID(id))).First(ctx)
	if ent.IsNotFound(err) {
		return StatusReport(c, 404)
	}

	err = db.User.Update().Where(user.HasReceiversWith(receiver.ID(id))).SetEmails(remove(userLoaded.Emails, email)).Exec(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function":   "delSome",
			"email":      email,
			"userLoaded": userLoaded,
			"error":      err,
		}).Errorln("Failed to delete an email address")
		return StatusReport(c, 500)
	}

	log.WithFields(log.Fields{
		"id":    c.Param("id"),
		"email": email,
	}).Infoln("Deleted an email address")

	return StatusReport(c, 200)
}

func listAll(c framework.Context) error {
	id := c.Param("id")
	client := c.Get("client").(string)

	resp, err := db.User.Query().WithReceivers().Where(user.HasReceiversWith(receiver.And(receiver.ID(id), receiver.HasConnectorWith(connector.ID(client))))).First(ctx)

	if ent.IsNotFound(err) {
		return c.JSON(http.StatusNotFound, []string{})
	} else if err != nil {
		log.WithFields(log.Fields{
			"function": "listAll",
		}).Error(err)
		return StatusReport(c, 500)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"emails": resp.Emails,
	})
}

func turnFwd(c framework.Context) error {
	id := c.Param("id")
	client := c.Get("client").(string)

	getUser, err := db.User.Query().Where(user.HasReceiversWith(receiver.And(receiver.ID(id), receiver.HasConnectorWith(connector.ID(client))))).First(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	err = db.User.Update().Where(user.HasReceiversWith(receiver.And(receiver.ID(id), receiver.HasConnectorWith(connector.ID(client))))).SetForward(!getUser.Forward).Exec(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	log.WithFields(log.Fields{
		"id":      id,
		"forward": !getUser.Forward,
	}).Infoln("Switched forward mode")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"forward": !getUser.Forward,
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
		url = "https://www." + osDomain
	}

	users, _ := db.User.Query().Where(user.Forward(true)).All(ctx)

	for _, u := range users {
		for _, r := range u.Edges.Receivers {
			finalRes := Result{
				Message: Message{
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
				return StatusReport(c, 400)
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
