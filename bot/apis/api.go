package apis

import (
	"gopkg.in/tucnak/telebot.v2"
)

// API wraps a Telegram bot for handling messages.
type API struct {
	*telebot.Bot
}

// WithBot wraps bot and returns an API instance for handling messages.
func WithBot(bot *telebot.Bot) *API {
	return &API{Bot: bot}
}
