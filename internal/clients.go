package main

import (
	"context"
	framework "github.com/labstack/echo/v4"
	"helium/ent"
	"helium/ent/connector"
	"net/http"
)

func createConnector(echoCtx framework.Context) error {
	c := new(ent.Connector)

	if err := echoCtx.Bind(c); err != nil {
		return echoCtx.String(http.StatusBadRequest, "Bad Request")
	}

	save, err := db.
		Connector.
		Create().
		SetID(c.ID).
		SetName(c.Name).
		SetSecret(c.Secret).
		SetURL(c.URL).
		Save(context.Background())
	if err != nil {
		return echoCtx.String(http.StatusInternalServerError, err.Error())
	}
	return echoCtx.JSON(http.StatusOK, save)
}

func fetchConnectors(c framework.Context) error {
	all, err := db.Connector.Query().Select("id", "name").All(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, all)
}

func getConnector(id string) string {
	first, err := db.Connector.Query().Where(connector.ID(id)).Select("url").First(context.Background())
	if err != nil {
		return ""
	}
	return first.URL
}
