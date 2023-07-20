package main

import (
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"helium/ent"
	"helium/ent/letter"
	"net/http"
)

func uploadHTML(html string, from string, to string) string {
	if len(html) == 0 || len(from) == 0 || len(to) == 0 {
		log.WithFields(log.Fields{"html": html, "from": from, "to": to}).Errorln("Did not upload letter due to lack of values")
		return ""
	}

	id := randSeq(8)
	for {
		count, err := db.Letter.Query().Where(letter.ID(id)).Count(ctx)
		if err != nil {
			break
		}
		if count == 0 {
			break
		}
		id = randSeq(8)
	}
	save, err := db.Letter.Create().SetID(id).SetHTML(html).SetFrom(from).SetTo(to).Save(ctx)
	if err != nil {
		return ""
	}

	return save.ID
}

func getRaw(c framework.Context) error {
	id := c.Param("id")

	l, err := db.Letter.Query().Where(letter.ID(id)).First(ctx)
	if ent.IsNotFound(err) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"id":   id,
			"html": "<html><head><title>404 Not Found</title></head><body><center><h1>404 Not Found</h1><hr>Atomic Emails</center></body>",
			"from": "entity_not_found@" + osDomain,
			"to":   "entity_not_found@" + osDomain,
		})
	}
	if err != nil {
		return c.String(400, http.StatusText(400))
	}

	return c.JSON(http.StatusOK, l)
}

func getHTML(c framework.Context) error {
	id := c.Param("id")

	l, err := db.Letter.Query().Where(letter.ID(id)).First(ctx)

	if ent.IsNotFound(err) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"id":   id,
			"html": "<html><head><title>404 Not Found</title></head><body><center><h1>404 Not Found</h1><hr>Atomic Emails</center></body>",
			"from": "entity_not_found@" + osDomain,
			"to":   "entity_not_found@" + osDomain,
		})
	}
	if err != nil {
		return c.String(400, http.StatusText(400))
	}

	return c.HTML(http.StatusOK, l.HTML)
}
