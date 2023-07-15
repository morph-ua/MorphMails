package main

import (
	"errors"
	"fmt"
	framework "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var rg = regexp.MustCompile(`(\r\n?|\n){2,}`)

func ParseAndSend(c framework.Context) error {
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

		var account Account
		result := db.Select("id", "forward", "paid", "times_received", "clients").Where("? = ANY(emails)", strings.ToLower(recipient)).First(&account)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			continue
		}

		if account.Forward {
			htmlRendered := "https://www.decline.live/preview/" + uploadHTML(html, from, rawRecipient)

			count, _ := strconv.Atoi(atc)
			var files []string
			if count > 0 {
				for i := 1; i <= count; i++ {
					file, err := c.FormFile(fmt.Sprintf("attachment-%d", i))
					if err != nil {
						return c.JSON(http.StatusBadRequest, HttpError{
							Status:  http.StatusBadRequest,
							Message: "Malformed file uploaded",
						})
					}
					fileBuf, err := file.Open()
					if err != nil {
						return c.JSON(http.StatusBadRequest, HttpError{
							Status:  http.StatusBadRequest,
							Message: "Malformed file uploaded",
						})
					}
					fileRead, err := io.ReadAll(fileBuf)
					if err != nil {
						return c.JSON(http.StatusBadRequest, HttpError{
							Status:  http.StatusBadRequest,
							Message: "Malformed file uploaded",
						})
					}
					var result CDNResponse
					_, err = reqClient.R().
						SetFileBytes("file", file.Filename, fileRead).
						SetHeader("Accept", "application/json").
						SetSuccessResult(&result).
						Post("https://cdn.lowt.live")
					if err != nil {
						return c.JSON(http.StatusBadRequest, HttpError{
							Status:  http.StatusBadRequest,
							Message: "Failed to upload one of the files to CDN",
						})
					}

					files = append(files, result.Message)
				}
			}
			finalRes := FinalResult{
				Message:     fmt.Sprintf(messageTemplate, from, rawRecipient, subject, text),
				RenderedURI: htmlRendered,
				ID:          account.ID,
				Files:       files,
			}

			log.WithFields(log.Fields{
				"recipients":    recipients,
				"ID":            account.ID,
				"clients":       account.Clients,
				"subject":       subject,
				"renderedEmail": htmlRendered,
			}).Infoln("Successfully parsed an email")

			for _, client := range account.Clients {
				_ = syncWithClients(finalRes, client, c)
			}
			return c.String(http.StatusOK, "OK")
		}
	}
	return c.JSON(http.StatusBadRequest, badRequestMessage)
}
