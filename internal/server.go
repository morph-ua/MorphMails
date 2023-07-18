package main

import (
	"context"
	"fmt"
	cron "github.com/go-co-op/gocron"
	request "github.com/imroc/req/v3"
	framework "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"helium/ent"
	"helium/ent/connector"
	"helium/ent/letter"
	"helium/ent/user"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var osSecret = os.Getenv("SECRET_KEY")
var osDB = os.Getenv("DATABASE_URL")
var osPort = os.Getenv("PORT")
var db *ent.Client
var req *request.Client

func timesReceivedNullification() {
	log.WithFields(log.Fields{
		"function": "timesReceivedNullification",
	}).Infoln("Running a TimesReceived nullification cronjob")
	_, err := db.User.Update().Where(user.Paid(false)).SetCounter(0).Save(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"function": "timesReceivedNullification",
			"error":    err,
		}).Fatalln("Failed to execute timesReceivedNullification")
	}
}

func letterNullification() {
	log.WithFields(log.Fields{
		"function": "letterNullification",
	}).Infoln("Running a letter nullification cronjob")
	_, err := db.Letter.Delete().Where(letter.CreatedAtLT(time.Now().AddDate(0, 0, -3))).Exec(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"function": "letterNullification",
			"error":    err,
		}).Fatalln("Failed to execute letterNullification")
	}
}

func runCronjobs() {
	s := cron.NewScheduler(time.UTC)
	if _, err := s.Every(1).Day().At("00:00").Do(timesReceivedNullification); err != nil {
		log.WithFields(log.Fields{
			"fatal":    true,
			"function": "runCronjobs",
			"error":    err,
		}).Fatalln("timesReceivedNullification CronJob failed to initialise")
	}

	if _, err := s.Every(3).Days().At("00:00").Do(letterNullification); err != nil {
		log.WithFields(log.Fields{
			"fatal":    true,
			"function": "runCronjobs",
			"error":    err,
		}).Fatalln("letterNullification CronJob failed to initialise")
	}

	s.StartAsync()
}

func checkV2Auth(next framework.HandlerFunc) framework.HandlerFunc {
	return func(c framework.Context) error {
		token := c.QueryParam("token")
		id, err := db.Connector.Query().Where(connector.Secret(token)).FirstID(context.Background())
		if err != nil {
			return err
		}
		if len(id) == 0 {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("client", id)
		return next(c)
	}
}

func checkSystemAuth(next framework.HandlerFunc) framework.HandlerFunc {
	return func(c framework.Context) error {
		token := c.QueryParam("token")
		if token != osSecret {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true, TimestampFormat: time.RFC822})
	log.SetOutput(os.Stdout)

	if len(osSecret) == 0 || len(osDB) == 0 {
		log.WithFields(log.Fields{
			"function": "init",
		}).Fatalln("Program failed to initialise. Required environment variables not found: `SECRET_KEY`, `DATABASE_URL`")
	}

	db, err := ent.Open("postgres", osDB)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "init",
		}).Fatalln("Failed to open database connection.")
	}

	defer func(db *ent.Client) {
		err := db.Close()
		if err != nil {
			log.WithFields(log.Fields{
				"function": "init (defer)",
			}).Fatalln("Failed to close database connection.")
		}
	}(db)

	req = request.C()

	if err := db.Schema.Create(context.Background()); err != nil {
		log.WithFields(log.Fields{
			"function": "init",
			"error":    err,
		}).Fatalln("Database migration was unsuccessful")
	}

	if len(osPort) == 0 {
		osPort = "8080"
	}
}

func main() {
	e := framework.New()
	e.HideBanner = true
	fmt.Println("Atomic Emails --> Helium V2")

	runCronjobs()

	e.Use(
		middleware.Gzip(),
		middleware.Recover(),
	)

	api := e.Group("/v2")

	api.Use(checkV2Auth)

	api.GET("/register/:id", registerNew)
	api.GET("/assign/:id", assignNew)
	api.GET("/forward/:id", turnFwd)
	api.GET("/delete/:id/:email", delSome)
	api.GET("/reset/:id", delAll)
	api.GET("/list/:id", listAll)

	system := e.Group("/sys")

	system.Use(checkSystemAuth)

	system.POST("/parse", ParseAndSend)
	system.POST("/announcement", sendAnnouncement)
	system.POST("/create/client", createConnector)

	e.GET("/html/:id", getHTML)
	e.GET("/data/:id", getRaw)
	e.GET("/clients", fetchConnectors)

	go func() {
		if err := e.Start(":" + osPort); err != nil && err != http.ErrServerClosed {
			log.Infoln("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
