package main

import (
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"helium/ent"
	"helium/ent/connector"
	"net/http"
)

func createConnector(context framework.Context) error {
	c := new(ent.Connector)

	if err := context.Bind(c); err != nil {
		return StatusReport(context, 400)
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
		return StatusReport(context, 500)
	}
	return context.JSON(200, save)
}

func fetchConnectors(c framework.Context) error {
	all, err := db.Connector.Query().Select("id", "name").All(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(200, all)
}

func getConnector(id string) string {
	first, err := db.Connector.Query().Where(connector.ID(id)).Select("url").First(ctx)
	if err != nil {
		return ""
	}
	return first.URL
}
