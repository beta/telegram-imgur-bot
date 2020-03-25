package apis

// Error replies for unexpected errors.
func (api *API) Error(req Request) {
	const errorMessage = `Sorry but there was some unexpected error. Please try again later.`
	api.Send(req.Sender(), errorMessage)
}

// Unsupported replies to all unsupported messages.
func (api *API) Unsupported(req Request) {
	const unsupportedMessage = `Sorry but I cannot handle that. Please send me an image (as a photo or file).`
	api.Send(req.Sender(), unsupportedMessage)
}
