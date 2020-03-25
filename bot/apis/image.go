package apis

import (
	"context"
	"fmt"
	"time"

	"github.com/beta/telegram-imgur-bot/bot/image"
	"github.com/beta/telegram-imgur-bot/bot/imgur"

	"gopkg.in/tucnak/telebot.v2"
)

// Photo handles messages with photos (compressed images).
func (api *API) Photo(m *Message) {
	if m.Photo == nil {
		api.Unsupported(m)
		return
	}
	if !m.Photo.InCloud() {
		api.Error(m)
		return
	}

	api.LogInfof(m, "[Photo] fileID=%s", m.Photo.FileID)

	image, err := api.uploadImageFromFileID(m.Photo.FileID, m.Caption)
	if err != nil {
		api.LogErrorf(m, "[Photo] %v", err)
		api.Error(m)
		return
	}

	api.LogInfof(m, "[Photo] image (fileID=%s) uploaded to Imgur, url=%s", m.Photo.FileID, image.URL)
	api.Reply(m.Message, image.URL, telebot.NoPreview)
}

// File handles messages with image files (uncompressed).
func (api *API) File(m *Message) {
	if m.Document == nil || !image.IsSupportedType(m.Document.MIME) {
		api.Unsupported(m)
		return
	}

	api.LogInfof(m, "[File] fileID=%s", m.Document.FileID)

	image, err := api.uploadImageFromFileID(m.Document.FileID, m.Caption)
	if err != nil {
		api.LogErrorf(m, "[File] %v", err)
		api.Error(m)
		return
	}

	api.LogInfof(m, "[File] image (fileID=%s) uploaded to Imgur, url=%s", m.Document.FileID, image.URL)
	api.Reply(m.Message, image.URL, telebot.NoPreview)
}

func (api *API) uploadImageFromFileID(fileID, caption string) (*imgur.Image, error) {
	// Get file URL.
	imageURL, err := api.FileURLByID(fileID)
	if err != nil {
		return nil, fmt.Errorf("error while querying image URL from file ID (%s): %v", fileID, err)
	}

	// Upload to Imgur.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	image, err := imgur.GetClient().UploadImage(ctx, imageURL, caption)
	if err != nil {
		return nil, fmt.Errorf("error while uploading image to Imgur: %v", err)
	}

	return image, nil
}
