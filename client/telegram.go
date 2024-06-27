package client

import (
	"log"
	"rivers-memo-cli/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.Envget("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	return bot
}
