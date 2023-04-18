package telegram

import (
	"log"
	"telegram-bot/pkg/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectURL     string
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, tr repository.TokenRepository, redirectURL string) *Bot {
	return &Bot{
		bot:             bot,
		pocketClient:    pocketClient,
		redirectURL:     redirectURL,
		tokenRepository: tr,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates, err := b.initUpdateChannel()
	if err != nil {
		return err
	}
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			// handle errors
			continue
		}
		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdateChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
