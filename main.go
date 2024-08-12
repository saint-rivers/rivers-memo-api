package main

import (
	"database/sql"
	"log"
	// "net"
	"net/http"
	"rivers-memo-cli/client"
	"rivers-memo-cli/config"
	// pb "rivers-memo-cli/generated/memo"
	"rivers-memo-cli/services/links"
	"rivers-memo-cli/services/tags"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	// "google.golang.org/grpc"
)

var exit = make(chan bool)

const PORT = ":5050"
const DIST = "ui"

func main() {
	db := client.InitDatabase()
	defer db.Close()

	e := echo.New()

	// enable REST API or not
	if config.Envget("REST_API_ENABLED") == "true" {
		e = configureRestApi(e, db)
		// configureGrpcServer()
	}

	// the production build serves the built html files from angular
	if config.Envget("EMBED_UI") == "true" {
		e.Static("/", "ui")
	}

	// this checks if telegram client should be started
	if config.Envget("TELEGRAM_ENABLED") == "true" {
		bot := client.InitTelegramBot()
		log.Printf("Telegram authorized on account %s", bot.Self.UserName)
		go listenTelegram(bot, db)
	}

	err := e.Start(PORT)
	if err != nil {
		e.Logger.Fatal(err)
		<-exit
		return
	}
	<-exit
}

func configureGrpcServer() {
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen on port 50051: %v", err)
	// }

	// s := grpc.NewServer()
	// server := &links.MemoServer{}
	// pb.RegisterMemoServiceServer(s, server)
	// log.Printf("gRPC server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}

func configureRestApi(e *echo.Echo, db *sql.DB) *echo.Echo {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200", "http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "up"})
	})
	links.ConfigRoutes(e, db)
	tags.ConfigRoutes(e, db)
	return e
}

func listenTelegram(bot *tgbotapi.BotAPI, db *sql.DB) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		// log.Printf("LOGGING: [%s]", update.ChannelPost.Text)
		err := links.ProcessChanncelMessage(db, &links.CreateLinkMemoParams{
			// ID:             update.ChannelPost.MessageID,
			ChannelMessage: update.ChannelPost.Text,
		})
		if err != nil {
			log.Println(err)
		}
	}
	exit <- true
}
