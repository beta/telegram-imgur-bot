package data

// Image represents an uploaded image.
type Image struct {
	ID                     int64  `db:"id"`
	TelegramUserID         int64  `db:"telegram_user_id"`
	TelegramReplyMessageID int64  `db:"telegram_reply_msg_id"`
	ImgurURL               string `db:"imgur_url"`
	ImgurDeleteHash        string `db:"imgur_delete_hash"`
}

const (
	mimeJPEG = "image/jpeg"
	mimePNG  = "image/png"
)

var supportedMIMEs = map[string]bool{
	mimeJPEG: true,
	mimePNG:  true,
}

// IsSupportedImageType returns whether the MIME is a supported one for images.
func IsSupportedImageType(mime string) bool {
	return supportedMIMEs[mime]
}
