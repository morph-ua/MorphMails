package main

import (
	"net/http"
	"regexp"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"helium/ent"
	"helium/ent/user"
)

var rg = regexp.MustCompile(`(\r\n?|\n){2,}`)

func unwrapDefaults(c echo.Context) unwrappedDefaults {
	recipients := strings.Split(c.FormValue("recipient"), ",")
	from := c.FormValue("from")
	subject := c.FormValue("subject")
	html := c.FormValue("stripped-html")
	text := rg.ReplaceAllString(c.FormValue("stripped-text"), "$1")

	switch {
	case len(subject) == 0:
		subject = "[No Subject]"
	case len(text) == 0:
		text = "[No Body]"
	case len(html) == 0:
		html = "[No Body]"
	}

	return unwrappedDefaults{
		Recipients: recipients,
		From:       from,
		Subject:    subject,
		HTML:       html,
		Text:       text,
	}
}

func parseAndSync(c echo.Context) error {
	values := unwrapDefaults(c)

	for _, recipient := range values.Recipients {
		split := strings.Split(recipient, "@")

		firstUser, err := db.User.
			Query().
			WithReceivers(func(query *ent.ReceiverQuery) {
				query.WithConnector()
			}).
			Select("forward", "paid", "counter").
			Where(func(selector *sql.Selector) {
				selector.Where(sqljson.ValueContains(user.FieldEmails, strings.Split(split[0], "+")[0]+"@"+split[1]))
			}).
			First(ctx)

		if ent.IsNotFound(err) {
			continue
		} else if err != nil {
			log.WithFields(log.Fields{
				"function": "parseAndSend",
			}).Error(err)
			return StatusReport(c, http.StatusInternalServerError)
		}

		if firstUser.Forward {
			go syncConnectors(
				result{
					Message: message{
						From:    values.From,
						To:      recipient,
						Subject: values.Subject,
						Text:    values.Text,
					},
					RenderedURI: "https://www.decline.live/preview/" +
						uploadHTML(values.HTML, values.From, recipient),
					Files: uploadFiles(c),
				},
				firstUser.Edges.Receivers,
			)

			return StatusReport(c, http.StatusOK)
		}
	}

	return StatusReport(c, http.StatusOK)
}
