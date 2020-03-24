package apis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/beta/imgur-bot/bot/imgur"

	"gopkg.in/tucnak/telebot.v2"
)

// Photo handles messages with photos (compressed images).
func (api *API) Photo(m *telebot.Message) {
	if m.Photo == nil {
		api.Unsupported(m)
		return
	}
	if !m.Photo.InCloud() {
		api.Error(m)
		return
	}

	log.Printf("%d [Photo] fileID=%s", m.ID, m.Photo.FileID)

	// Get image URL.
	imageURL, err := api.FileURLByID(m.Photo.FileID)
	if err != nil {
		log.New(os.Stderr, "", log.LstdFlags).Printf("%d [Photo] error while querying image URL from its ID (%s): %v", m.ID, m.Photo.FileID, err)
		api.Error(m)
		return
	}

	// Upload to Imgur.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	image, err := imgur.GetClient().UploadImage(ctx, imageURL, m.Caption)
	if err != nil {
		log.New(os.Stderr, "", log.LstdFlags).Printf("%d [Photo] error while uploading image to Imgur: %v", m.ID, err)
		api.Error(m)
		return
	}

	log.Printf("%d [Photo] image uploaded to Imgur, URL: %s", m.ID, image.URL)
	api.Reply(m, image.URL, telebot.NoPreview)
}
