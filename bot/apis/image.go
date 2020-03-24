package apis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beta/imgur-bot/bot/image"
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

	image, err := api.uploadImageFromFileID(m.Photo.FileID, m.Caption)
	if err != nil {
		log.New(os.Stderr, "", log.LstdFlags).Printf("%d [Photo] %v", m.ID, err)
		api.Error(m)
		return
	}

	api.Reply(m, image.URL, telebot.NoPreview)
}

// File handles messages with image files (uncompressed).
func (api *API) File(m *telebot.Message) {
	if m.Document == nil || !image.IsSupportedType(m.Document.MIME) {
		api.Unsupported(m)
		return
	}

	log.Printf("%d [File] fileID=%s", m.ID, m.Document.FileID)

	image, err := api.uploadImageFromFileID(m.Document.FileID, m.Caption)
	if err != nil {
		log.New(os.Stderr, "", log.LstdFlags).Printf("%d [File] %v", m.ID, err)
		api.Error(m)
		return
	}

	api.Reply(m, image.URL, telebot.NoPreview)
}

func (api *API) uploadImageFromFileID(fileID, caption string) (*imgur.Image, error) {
	// Get file URL.
	imageURL, err := api.FileURLByID(fileID)
	if err != nil {
		return nil, fmt.Errorf("error while querying image URL from file ID (%s): %v", fileID, err)
	}

	// Upload to Imgur.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	image, err := imgur.GetClient().UploadImage(ctx, imageURL, caption)
	if err != nil {
		return nil, fmt.Errorf("error while uploading image to Imgur: %v", err)
	}

	return image, nil
}
