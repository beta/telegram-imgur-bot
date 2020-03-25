package middlewares

import (
	"strings"

	"github.com/beta/telegram-imgur-bot/bot/apis"
	"github.com/beta/telegram-imgur-bot/bot/image"

	"gopkg.in/tucnak/telebot.v2"
)

// Logger outputs logs for all messages received.
func Logger(bot *telebot.Bot) func(update *telebot.Update) bool {
	return func(update *telebot.Update) bool {
		if update == nil {
			return false
		}

		api := apis.WithBot(bot)

		switch {
		case update.Message != nil:
			m := &apis.Message{Message: update.Message}
			api.LogInfof(m, "[Message] updateID=%d, sender=%s, content=%s, caption=%s, hasImage=%v, hasImageFile=%v",
				update.ID, getSenderName(m.Sender()), m.Text, m.Caption, (m.Photo != nil),
				(m.Document != nil && image.IsSupportedType(m.Document.MIME)))

		case update.Callback != nil:
			cb := &apis.Callback{Callback: update.Callback}
			api.LogInfof(cb, "[Callback] updateID=%d, sender=%s, messageID=%d, data=%s", update.ID, getSenderName(cb.Sender()), cb.MessageID, cb.Data)
		}

		return true
	}
}

func getSenderName(sender *telebot.User) string {
	if sender == nil {
		return ""
	}

	parts := make([]string, 0, 3)
	if sender.FirstName != "" {
		parts = append(parts, sender.FirstName)
	}
	if sender.LastName != "" {
		parts = append(parts, sender.LastName)
	}
	if sender.Username != "" {
		parts = append(parts, "(@"+sender.Username+")")
	}
	return strings.Join(parts, " ")
}
