package apis

import (
	"context"
	"fmt"
	"time"

	"github.com/beta/telegram-imgur-bot/bot/data"
	"github.com/beta/telegram-imgur-bot/bot/db"
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

	image, err := api.uploadTelegramImageFile(m, m.Photo.File)
	if err != nil {
		api.LogErrorf(m, "[Photo] %v", err)
		api.Error(m)
		return
	}

	api.LogInfof(m, "[Photo] image (fileID=%s) uploaded to Imgur, url=%s", m.Photo.FileID, image.ImgurURL)
	api.Reply(m.Message, image.ImgurURL, telebot.NoPreview)
}

// File handles messages with image files (uncompressed).
func (api *API) File(m *Message) {
	if m.Document == nil || !data.IsSupportedImageType(m.Document.MIME) {
		api.Unsupported(m)
		return
	}

	api.LogInfof(m, "[File] fileID=%s", m.Document.FileID)

	image, err := api.uploadTelegramImageFile(m, m.Document.File)
	if err != nil {
		api.LogErrorf(m, "[File] %v", err)
		api.Error(m)
		return
	}

	api.LogInfof(m, "[File] image (fileID=%s) uploaded to Imgur, url=%s", m.Document.FileID, image.ImgurURL)
	api.Reply(m.Message, image.ImgurURL, telebot.NoPreview)
}

func (api *API) uploadTelegramImageFile(m *Message, file telebot.File) (*data.Image, error) {
	// Get file URL.
	imageURL, err := api.FileURLByID(file.FileID)
	if err != nil {
		return nil, fmt.Errorf("error while querying image URL from file ID (%s): %v", file.FileID, err)
	}

	// Upload to Imgur.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uploaded, err := imgur.GetClient().UploadImage(ctx, imageURL, m.Caption)
	if err != nil {
		return nil, fmt.Errorf("error while uploading image to Imgur: %v", err)
	}

	// Save to DB.
	image := &data.Image{
		TelegramUserID:  int64(m.Sender().ID),
		ImgurURL:        uploaded.URL,
		ImgurDeleteHash: uploaded.DeleteHash,
	}
	inserted, err := db.Image.Insert(image)
	if err != nil {
		return nil, fmt.Errorf("error while saving image into DB: %v", err)
	}

	return inserted, nil
}
