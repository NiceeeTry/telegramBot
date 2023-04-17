package main

import (
	"log"
	"telegram-bot/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6245099511:AAFiKBS-4hn3cCDUdvYU4xODAC4nsuhv2Ds")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	pocketClient, err := pocket.NewClient("106869-dca7518eb58b445d8d67fce")
	if err != nil {
		log.Fatal(err)
	}
	telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
