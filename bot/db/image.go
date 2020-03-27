package db

import (
	"github.com/beta/telegram-imgur-bot/bot/data"
)

// Image is the entry point for image DB APIs.
var Image = new(image)

// Image DB API implementation.
type image struct{}

// Insert inserts an image to DB and returns the inserted image.
func (*image) Insert(image *data.Image) (*data.Image, error) {
	const query = `INSERT INTO "images"
			("telegram_user_id", "telegram_reply_msg_id", "imgur_url", "imgur_delete_hash")
		VALUES ($1, $2, $3, $4)
		RETURNING "id", "telegram_user_id", "telegram_reply_msg_id", "imgur_url", "imgur_delete_hash"`

	ret := new(data.Image)
	err := db.QueryRow(query, image.TelegramUserID, image.TelegramReplyMessageID, image.ImgurURL, image.ImgurDeleteHash).
		Scan(&ret.ID, &ret.TelegramUserID, &ret.TelegramReplyMessageID, &ret.ImgurURL, &ret.ImgurDeleteHash)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Query queries an image with its ID from DB.
func (*image) Query(id int64) (*data.Image, error) {
	const query = `SELECT "id", "telegram_user_id", "telegram_reply_msg_id", "imgur_url", "imgur_delete_hash" FROM "images" WHERE "id" = $1`

	image := new(data.Image)
	err := db.QueryRow(query, id).
		Scan(&image.ID, &image.TelegramUserID, &image.TelegramReplyMessageID, &image.ImgurURL, &image.ImgurDeleteHash)
	if err != nil {
		return nil, err
	}
	return image, nil
}

// Delete deletes an image.
func (*image) Delete(id int64) error {
	const query = `DELETE FROM "images" WHERE "id" = $1`

	if _, err := db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
