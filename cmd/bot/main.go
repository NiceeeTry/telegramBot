package main

import (
	"log"
	"telegram-bot/pkg/repository"
	"telegram-bot/pkg/repository/boltdb"
	"telegram-bot/pkg/telegram"

	"github.com/boltdb/bolt"
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
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	tokenRepository := boltdb.NewTokenRepository(db)
	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")

	// 46:40
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return db, nil
}
