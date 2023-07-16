package main

import (
	"errors"
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func createClient(context framework.Context) error {
	c := new(Client)

	if err := context.Bind(c); err != nil {
		return context.String(http.StatusBadRequest, "Bad Request")
	}

	client := Client{
		ID:     c.ID,
		Name:   c.Name,
		URL:    c.URL,
		Secret: c.Secret,
	}

	db.Create(client)

	return context.JSON(http.StatusOK, client)
}

func fetchClients(c framework.Context) error {
	var results []map[string]interface{}

	db.Model(&Client{}).Select("id", "name").Find(&results)

	return c.JSON(http.StatusOK, results)
}

func getClient(id string) string {
	var client Client
	result := db.Model(&Client{}).Where("id = ?", id).First(&client)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Errorf("The client %s wasn't found in DB.\n", id)
		return ""
	}
	return client.URL
}
