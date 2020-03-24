package apis

import (
	"fmt"

	"gopkg.in/tucnak/telebot.v2"
)

const helloMessage = `Hi %s :) I can help you upload images to Imgur. Send me an image directly.`

// Hello says hello to the user.
func (api *API) Hello(m *telebot.Message) {
	api.Send(m.Sender, fmt.Sprintf(helloMessage, m.Sender.FirstName))
}
