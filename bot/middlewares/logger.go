package middlewares

import (
	"log"
	"strings"

	"github.com/beta/imgur-bot/bot/image"
	"gopkg.in/tucnak/telebot.v2"
)

// Logger outputs logs for all messages received.
func Logger(update *telebot.Update) bool {
	if update == nil {
		return false
	}

	switch {
	case update.Message != nil:
		m := update.Message
		log.Printf("[Message] updateID=%d, messageID=%d, sender=%s, content=%s, hasImage=%v, hasImageFile=%v",
			update.ID, m.ID, getSenderName(m.Sender), m.Text, (m.Photo != nil),
			(m.Document != nil && image.IsSupportedType(m.Document.MIME)))
	}

	return true
}

func getSenderName(sender *telebot.User) string {
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
