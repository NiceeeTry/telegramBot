package telegram

import (
	"context"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

const (
	commandStart           = "start"
	replyStartTemplate     = "Hi, It is necessery to give me access of your Pocket account to store the links. To do it go through the link:\n%s"
	replyAlreadyAuthorized = "You are already authorized. Just send a link, I will save it"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	// log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, "The link is saved")
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		msg.Text = "The link is not valid"
		_, err := b.bot.Send(msg)
		return err
	}
	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		msg.Text = "You are not authorized. Use the /start command"
		_, err := b.bot.Send(msg)
		return err
	}
	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		msg.Text = "The link wasnt saved, try again"
		_, err := b.bot.Send(msg)
		return err
	}
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAuthorized)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don`t know such a command(")
	_, err := b.bot.Send(msg)
	return err
}
