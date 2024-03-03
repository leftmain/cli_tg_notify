package telegram

import (
	"send2telegram/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(config config.SendConfig, msg string) error {
	bot, err := tgbotapi.NewBotAPI(config.TelegramApiToken)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(config.ChatId, msg)
	if _, err := bot.Send(message); err != nil {
		return err
	}

	return nil
}
