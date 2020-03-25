package apis

import (
	"fmt"
)

// Hello says hello to the user.
func (api *API) Hello(m *Message) {
	const helloMessage = `Hi %s :) I can help you upload images to Imgur. Send me an image (as a photo or file).`
	api.Send(m.Sender(), fmt.Sprintf(helloMessage, m.Sender().FirstName))
}
