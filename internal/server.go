package main

import (
	"context"
	"errors"
	"fmt"
	cron "github.com/go-co-op/gocron"
	request "github.com/imroc/req/v3"
	framework "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var osSecret = os.Getenv("SECRET_KEY")
var osDB = os.Getenv("DATABASE_URL")
var osPort = os.Getenv("PORT")
var db *gorm.DB
var req *request.Client

func timesReceivedNullification(DB *gorm.DB) {
	log.WithFields(log.Fields{
		"function": "timesReceivedNullification",
	}).Infoln("Running a TimesReceived nullification cronjob")
	DB.Model(&Account{}).Where("paid = ?", false).Update("times_received", 0)
}

func letterNullification(DB *gorm.DB) {
	log.WithFields(log.Fields{
		"function": "letterNullification",
	}).Infoln("Running a letter nullification cronjob")
	DB.Where("created_at < ?", time.Now().UTC().Add(-1*24*time.Duration(7)*time.Hour)).Delete(&Letter{})
}

func runCronJob(DB *gorm.DB) {
	s := cron.NewScheduler(time.UTC)
	if _, err := s.Every(1).Day().At("00:00").Do(timesReceivedNullification, DB); err != nil {
		log.WithFields(log.Fields{
			"fatal":    true,
			"function": "runCronJob",
			"error":    err,
		}).Fatalln("timesReceivedNullification CronJob failed to initialise")
	}

	if _, err := s.Every(3).Days().At("00:00").Do(letterNullification, DB); err != nil {
		log.WithFields(log.Fields{
			"fatal":    true,
			"function": "runCronJob",
			"error":    err,
		}).Fatalln("letterNullification CronJob failed to initialise")
	}

	s.StartAsync()
}

func CheckV2Auth(next framework.HandlerFunc) framework.HandlerFunc {
	return func(c framework.Context) error {
		token := c.QueryParam("token")
		var result Client
		response := db.Where("secret = ?", token).Select("id").First(&result)
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		c.Set("client", result.ID)

		return next(c)
	}
}

func CheckSystemAuth(next framework.HandlerFunc) framework.HandlerFunc {
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
			"fatal":    true,
			"function": "init",
		}).Fatalln("Program failed to initialise. Required environment variables not found: `SECRET_KEY`, `DATABASE_URL`")
	}

	db = connect()
	req = request.C()

	if err := db.AutoMigrate(&Account{}, &Letter{}, &Client{}); err != nil {
		log.WithFields(log.Fields{
			"fatal":    true,
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

	runCronJob(db)

	e.Use(
		middleware.Gzip(),
		middleware.Recover(),
	)

	api := e.Group("/v2")

	api.Use(CheckV2Auth)

	api.GET("/register/:id", registerNew)
	api.GET("/assign/:id", assignNew)
	api.GET("/forward/:id", turnFwd)
	api.GET("/delete/:id/:email", delSome)
	api.GET("/reset/:id", delAll)
	api.GET("/list/:id", listAll)

	system := e.Group("/sys")

	system.Use(CheckSystemAuth)

	system.POST("/parse", ParseAndSend)
	system.POST("/announcement", sendAnnouncement)
	system.POST("/create/client", createClient)

	e.GET("/html/:id", getHTML)
	e.GET("/data/:id", getRaw)
	e.GET("/clients", fetchClients)

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
