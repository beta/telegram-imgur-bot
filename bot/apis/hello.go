package apis

import (
	"fmt"

	"gopkg.in/tucnak/telebot.v2"
)

// Hello says hello to the user.
func (api *API) Hello(m *telebot.Message) {
	const helloMessage = `Hi %s :) I can help you upload images to Imgur. Send me an image (as a photo or file).`
	api.Send(m.Sender, fmt.Sprintf(helloMessage, m.Sender.FirstName))
}
