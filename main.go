package main

import (
	"database/sql"
	"log"
	"net/http"
	"rivers-memo-cli/client"
	"rivers-memo-cli/services/links"
	"rivers-memo-cli/services/tags"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

var exit = make(chan bool)

const PORT = ":5050"

func main() {
	bot := client.InitTelegramBot()
	db := client.InitDatabase()
	defer db.Close()
	log.Printf("Telegram authorized on account %s", bot.Self.UserName)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200", "http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "up"})
	})
	links.ConfigRoutes(e, db)
	tags.ConfigRoutes(e, db)

	go listenTelegram(bot, db)
	err := e.Start(PORT)
	if err != nil {
		e.Logger.Fatal(err)
		<-exit
		return
	}
	<-exit
}

func listenTelegram(bot *tgbotapi.BotAPI, db *sql.DB) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		// log.Printf("LOGGING: [%s]", update.ChannelPost.Text)
		err := links.ProcessChanncelMessage(db, &links.CreateLinkMemoParams{
			ID:             update.ChannelPost.MessageID,
			ChannelMessage: update.ChannelPost.Text,
		})
		if err != nil {
			log.Println(err)
		}
	}
	exit <- true
}
