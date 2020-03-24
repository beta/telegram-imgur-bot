package apis

import (
	"gopkg.in/tucnak/telebot.v2"
)

// Hello says hello to the user.
func Hello(bot *telebot.Bot) interface{} {
	return func(m *telebot.Message) {
		bot.Send(m.Sender, "Hello world")
	}
}
