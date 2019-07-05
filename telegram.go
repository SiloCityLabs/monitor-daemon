package main

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Send a telegram message or error
func telegram(body string) error {
	bot, err := tgbotapi.NewBotAPI(Settings.Telegram.Apikey)
	if err != nil {
		return errors.New("Failed to connect telegram bot: " + err.Error())
	}

	// How to get chat id: https://api.telegram.org/bot732500886:AAHC4MOmURRF5FlJTbvWZ2GlIG4pLnkZSKU/getUpdates
	msg := tgbotapi.NewMessage(Settings.Telegram.Chatid, body)
	msg.ParseMode = "markdown"

	if _, err := bot.Send(msg); err != nil {
		return errors.New("Failed to send telegram message: " + err.Error())
	}

	return nil
}
