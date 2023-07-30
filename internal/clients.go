package main

import (
	"morph_mails/ent"
	"net/http"

	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func createConnector(context framework.Context) error {
	c := new(ent.Connector)

	if err := context.Bind(c); err != nil {
		return StatusReport(context, http.StatusBadRequest)
	}

	save, err := db.
		Connector.
		Create().
		SetID(c.ID).
		SetName(c.Name).
		SetSecret(c.Secret).
		SetURL(c.URL).
		Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error":         err,
			"function":      "createConnector",
			"connectorData": c,
		})

		return StatusReport(context, http.StatusInternalServerError)
	}

	return context.JSON(http.StatusOK, save)
}

func fetchConnectors(c framework.Context) error {
	all, err := db.Connector.Query().Select("id", "name").All(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, all)
}
