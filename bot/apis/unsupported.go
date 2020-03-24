package apis

import (
	"gopkg.in/tucnak/telebot.v2"
)

// Unsupported replies to all unsupported messages.
func (api *API) Unsupported(m *telebot.Message) {
	const unsupportedMessage = `Sorry but I cannot handle that. Please send me an image (as a photo or file).`

	api.Send(m.Sender, unsupportedMessage)
}
