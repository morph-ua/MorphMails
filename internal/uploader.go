package main

import (
	"bytes"
	"morph_mails/ent"
	"morph_mails/ent/letter"
	"io"
	"mime/multipart"
	"net/http"

	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func uploadFiles(c framework.Context) []string {
	form, _ := c.MultipartForm()
	var links []string

	for _, files := range form.File {
		for _, file := range files {
			src := func() multipart.File {
				src, err := file.Open()
				if err != nil {
					return nil
				}
				defer src.Close()

				return src
			}()

			var buf bytes.Buffer
			if _, err := io.Copy(&buf, src); err != nil {
				return nil
			}

			var result map[string]string
			_, _ = req.R().
				SetFileBytes("file", file.Filename, buf.Bytes()).
				SetHeader("Accept", "application/json").
				SetSuccessResult(&result).
				Post("https://cdn.lowt.live")

			links = append(links, result["message"])
		}
	}

	return links
}

func uploadHTML(html string, from string, to string) string {
	if len(html) == 0 || len(from) == 0 || len(to) == 0 {
		log.WithFields(log.Fields{
			"html": html,
			"from": from,
			"to":   to,
		}).Errorln("Did not upload letter due to lack of values")

		return ""
	}

	id, err := sid.Generate()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "uploadHTML(sid.Generate)",
		}).Errorln(err)

		return ""
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
		return c.JSON(http.StatusNotFound, map[string]any{
			"id": id,
			"html": "<html><head><title>404 Not Found</title></head>" +
				"<body><center><h1>404 Not Found</h1><hr>Morph Mails</center></body>",
			"from": "entity_not_found@" + osDomain,
			"to":   "entity_not_found@" + osDomain,
		})
	}

	if err != nil {
		return StatusReport(c, http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, l)
}

func getHTML(c framework.Context) error {
	id := c.Param("id")

	l, err := db.Letter.Query().Where(letter.ID(id)).First(ctx)

	if ent.IsNotFound(err) {
		return c.HTML(http.StatusNotFound, "<html><head><title>404 Not Found</title></head>"+
			"<body><center><h1>404 Not Found</h1><hr>Morph Mails</center></body>")
	}

	if err != nil {
		return StatusReport(c, http.StatusBadRequest)
	}

	return c.HTML(http.StatusOK, l.HTML)
}
