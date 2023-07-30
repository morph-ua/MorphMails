package main

import (
	"helium/ent"
	"helium/ent/user"
	"net/http"
	"strconv"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func noMiscRenderPlugin() md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: []string{"img", "svg", "script"},
				Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
					content = ""

					return &content
				},
			},
		}
	}
}

func convert(html string) string {
	converter := md.NewConverter("", true, nil)
	converter.Use(noMiscRenderPlugin())

	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "convert",
		}).Error(err)
		return "❌ **Failed to parse Markdown from Letter**"
	}

	return markdown
}

func unwrapDefaults(c echo.Context) unwrappedDefaults {
	recipients := strings.Split(c.FormValue("recipient"), ",")
	from := c.FormValue("from")
	subject := c.FormValue("subject")
	html := c.FormValue("stripped-html")
	text := convert(html)
	atcStr := c.FormValue("attachment-count")
	switch {
	case len(subject) == 0:
		subject = "[No Subject]"
	case len(html) == 0:
		html = "[No Body]"
	case len(atcStr) == 0:
		atcStr = "0"
	}

	atc, _ := strconv.Atoi(atcStr)

	return unwrappedDefaults{
		Recipients: recipients,
		From:       from,
		Subject:    subject,
		HTML:       html,
		Text:       text,
		ATC:        atc,
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

		log.Println(firstUser, firstUser.Edges.Receivers)

		if ent.IsNotFound(err) {
			continue
		} else if err != nil {
			log.WithFields(log.Fields{
				"function": "parseAndSend",
			}).Error(err)
			return StatusReport(c, http.StatusInternalServerError)
		}

		if firstUser.Forward {
			var files []string
			if values.ATC == 0 {
				files = []string{}
			} else {
				files = uploadFiles(c)
			}

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
					Files: files,
				},
				firstUser.Edges.Receivers,
			)

			return StatusReport(c, http.StatusOK)
		}
	}

	return StatusReport(c, http.StatusOK)
}
