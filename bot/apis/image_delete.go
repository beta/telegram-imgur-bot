package apis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/beta/telegram-imgur-bot/bot/imgur"

	"github.com/beta/telegram-imgur-bot/bot/data"
	"github.com/beta/telegram-imgur-bot/bot/db"
	"gopkg.in/tucnak/telebot.v2"
)

const (
	errNoDeletePermission = `You have no permission to delete this image.`

	msgDeleteConfirmation = `Do you really want to delete this image? *This CANNOT be undone.*`
	msgDeleted            = `_Image deleted._`

	buttonDeleteCancel  = `🔙 No, go back`
	buttonDeleteConfirm = `🗑 Yes, delete it`
)

func (api *API) handleDelete(cb *Callback, data string) {
	// Query image.
	image, err := queryImageFromIDString(data)
	if err != nil {
		api.LogErrorf(cb, "[handleDelete] error while querying image from callback data [%s]: %v", data, err)
		api.ErrorCallback(cb)
		return
	}

	// Validate user permission.
	if image.TelegramUserID != int64(cb.Sender().ID) {
		api.LogErrorf(cb, "[handleDelete] image sender ID=%d, callback sender ID=%d, not allowed to delete image", image.TelegramUserID, cb.Sender().ID)
		api.Respond(cb.Callback, &telebot.CallbackResponse{
			Text:      errNoDeletePermission,
			ShowAlert: true,
		})
		return
	}

	api.Respond(cb.Callback)
	api.Edit(cb.Message, getImageMessage(image, msgDeleteConfirmation),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeMarkdown,
			DisableWebPagePreview: true,
		},
		&telebot.ReplyMarkup{
			InlineKeyboard: api.generateDeleteConfirmInlineKeyboard(image),
		},
	)
}

// Returns message content for replying to user.
func getImageMessage(image *data.Image, extra string) string {
	if len(extra) <= 0 {
		return image.ImgurURL
	}

	return fmt.Sprintf("%s\n\n%s", image.ImgurURL, extra)
}

// Queries an image from DB with ID as a string.
func queryImageFromIDString(idStr string) (*data.Image, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error while parsing image ID from string: %v", err)
	}

	return db.Image.Query(id)
}

func (api *API) generateDeleteConfirmInlineKeyboard(image *data.Image) [][]telebot.InlineButton {
	// Cancel button.
	cancelButton := telebot.InlineButton{
		Unique: actionDeleteCancel,
		Text:   buttonDeleteCancel,
		Data:   fmt.Sprint(image.ID),
	}

	// Confirm button.
	confirmButton := telebot.InlineButton{
		Unique: actionDeleteConfirm,
		Text:   buttonDeleteConfirm,
		Data:   fmt.Sprint(image.ID),
	}

	return [][]telebot.InlineButton{
		{cancelButton, confirmButton},
	}
}

func (api *API) handleDeleteConfirm(cb *Callback, data string) {
	// Query image.
	image, err := queryImageFromIDString(data)
	if err != nil {
		api.LogErrorf(cb, "[handleDeleteConfirm] error while querying image from callback data [%s]: %v", data, err)
		api.ErrorCallback(cb)
		return
	}

	// Validate user permission.
	if image.TelegramUserID != int64(cb.Sender().ID) {
		api.LogErrorf(cb, "[handleDeleteConfirm] image sender ID=%d, callback sender ID=%d, not allowed to delete image", image.TelegramUserID, cb.Sender().ID)
		api.Respond(cb.Callback, &telebot.CallbackResponse{
			Text:      errNoDeletePermission,
			ShowAlert: true,
		})
		return
	}

	// Delete image.
	if err := deleteImage(image); err != nil {
		api.LogErrorf(cb, "[handleDeleteConfirm] error while deleting image: %v", err)
		api.ErrorCallback(cb)
		return
	}

	api.LogInfof(cb, "[handleDeleteConfirm] image deleted, ID=%d", image.ID)
	api.Respond(cb.Callback)
	api.Edit(cb.Message, msgDeleted,
		&telebot.SendOptions{
			ParseMode: telebot.ModeMarkdown,
		},
	)
}

func deleteImage(image *data.Image) error {
	// Delete from Imgur.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := imgur.GetClient().DeleteImage(ctx, image.ImgurDeleteHash); err != nil {
		return fmt.Errorf("error while deleting image from Imgur: %v", err)
	}

	// Delete from DB.
	if err := db.Image.Delete(image.ID); err != nil {
		return fmt.Errorf("error while deleting image from DB: %v", err)
	}

	return nil
}

func (api *API) handleDeleteCancel(cb *Callback, data string) {
	// Query image.
	image, err := queryImageFromIDString(data)
	if err != nil {
		api.LogErrorf(cb, "[handleDeleteCancel] error while querying image from callback data [%s]: %v", data, err)
		api.ErrorCallback(cb)
		return
	}

	api.Respond(cb.Callback)
	api.Edit(cb.Message, getImageMessage(image, ""),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeMarkdown,
			DisableWebPagePreview: true,
		},
		&telebot.ReplyMarkup{
			InlineKeyboard: api.generateImageInlineKeyboard(image),
		},
	)
}
