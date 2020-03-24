package apis

import (
	"gopkg.in/tucnak/telebot.v2"
)

// Error replies for unexpected errors.
func (api *API) Error(m *telebot.Message) {
	const errorMessage = `Sorry but there was some unexpected error. Please try again later.`
	api.Send(m.Sender, errorMessage)
}

// Unsupported replies to all unsupported messages.
func (api *API) Unsupported(m *telebot.Message) {
	const unsupportedMessage = `Sorry but I cannot handle that. Please send me an image (as a photo or file).`

	api.Send(m.Sender, unsupportedMessage)
}
