package main

import (
	"log"
	"rivers-memo-cli/client"
	"rivers-memo-cli/services/links"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	// "github.com/alexflint/go-arg"

	_ "github.com/lib/pq"
)

func main() {
	bot := client.InitTelegramBot()
	db := client.InitDatabase()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// log.Printf("LOGGING: [%s]", update.ChannelPost.Text)
		err := links.CreateLinkMemo(db, &links.CreateLinkMemoParams{
			ID:   update.ChannelPost.MessageID,
			Link: update.ChannelPost.Text,
		})
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("running")
}
