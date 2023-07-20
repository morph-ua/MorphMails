package main

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"helium/ent"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var rg = regexp.MustCompile(`(\r\n?|\n){2,}`)

func ParseAndSend(c echo.Context) error {
	recipients := strings.Split(c.FormValue("recipient"), ",")
	from := c.FormValue("from")
	subject := c.FormValue("subject")
	html := c.FormValue("stripped-html")
	text := rg.ReplaceAllString(c.FormValue("stripped-text"), "$1")

	atc := c.FormValue("attachment-count")

	switch {
	case len(subject) == 0:
		subject = "[No Subject]"
	case len(text) == 0:
		text = "[No Body]"
	case len(html) == 0:
		html = "[No Body]"
	case len(atc) == 0:
		atc = "0"
	}

	for _, recipient := range recipients {
		var rawRecipient = recipient

		if strings.Contains(recipient, "+") {
			split := strings.Split(recipient, "@")
			recipient = strings.Split(split[0], "+")[0] + "@" + split[1]
		}

		user, err := db.User.Query().
			Select("forward", "paid", "counter", "receivers").
			Where(func(selector *sql.Selector) {
				selector.Where(sql.Contains("emails", recipient))
			}).
			WithReceivers(func(query *ent.ReceiverQuery) {
				query.WithConnector()
			}).
			First(ctx)

		if ent.IsNotFound(err) {
			continue
		} else if err != nil {
			return StatusReport(c, 500)
		}

		if user.Forward {
			htmlRendered := "https://www.decline.live/preview/" + uploadHTML(html, from, rawRecipient)

			count, _ := strconv.Atoi(atc)
			var files []string
			if count > 0 {
				for i := 1; i <= count; i++ {
					var fileContent []byte
					var file *multipart.FileHeader
					if err := func() error {
						file, err = c.FormFile(fmt.Sprintf("attachment-%d", i))
						if err != nil {
							return err
						}
						fileBuf, err := file.Open()
						if err != nil {
							return err
						}
						fileContent, err = io.ReadAll(fileBuf)
						if err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return StatusReport(c, 400)
					}
					var result FileUploader
					_, err = req.R().
						SetFileBytes("file", file.Filename, fileContent).
						SetHeader("Accept", "application/json").
						SetSuccessResult(&result).
						Post("https://cdn.lowt.live")
					if err != nil {
						return c.JSON(http.StatusBadRequest, Error{
							Status:  http.StatusBadRequest,
							Message: "Failed to upload one of the files to CDN",
						})
					}

					files = append(files, result.Message)
				}
			}

			log.WithFields(log.Fields{
				"recipients":    recipients,
				"ID":            user.ID,
				"subject":       subject,
				"renderedEmail": htmlRendered,
			}).Infoln("Successfully parsed an email")

			for _, receiver := range user.Edges.Receivers {
				result := Result{
					Message: Message{
						From: from,
						To:   recipient,
						Text: text,
					},
					RenderedURI: htmlRendered,
					ID:          receiver.ID,
					Files:       files,
				}

				receiver := receiver
				go syncConnectors(result, receiver.Edges.Connector.URL)
			}

			return StatusReport(c, 200)
		}
	}
	return StatusReport(c, 200)
}
