package apis

import (
	"gopkg.in/tucnak/telebot.v2"
)

const errorMessage = `Sorry but there was some unexpected error. Please try again later.`

// Error replies for unexpected errors.
func (api *API) Error(req Request) {
	api.Send(req.Sender(), errorMessage)
}

// ErrorCallback responds cb with unexpected error.
func (api *API) ErrorCallback(cb *Callback) {
	api.Respond(cb.Callback, &telebot.CallbackResponse{
		Text:      errorMessage,
		ShowAlert: true,
	})
}

// Unsupported replies to all unsupported messages.
func (api *API) Unsupported(req Request) {
	const unsupportedMessage = `Sorry but I cannot handle that. Please send me an image (as a photo or file).`
	api.Send(req.Sender(), unsupportedMessage)
}

// UnsupportedCallback responds cb with unsupported error.
func (api *API) UnsupportedCallback(cb *Callback) {
	const unsupportedCallback = `Sorry but I cannot handle that.`
	api.Respond(cb.Callback, &telebot.CallbackResponse{
		Text:      unsupportedCallback,
		ShowAlert: true,
	})
}
