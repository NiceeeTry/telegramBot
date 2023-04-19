package telegram

import (
	"context"
	"fmt"
	"telegram-bot/pkg/repository"
)

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID)
	requestTopken, err := b.pocketClient.GetRequestToken(context.Background(), b.redirectURL)
	if err != nil {
		return "", err
	}
	if err := b.tokenRepository.Save(chatID, requestTopken, repository.RequestTokens); err != nil {
		return "", err
	}
	return b.pocketClient.GetAuthorizationURL(requestTopken, redirectURL)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}
