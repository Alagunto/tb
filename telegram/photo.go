package telegram

import "github.com/alagunto/tb/files"

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
// In the Telegram Bot API, photos are represented as arrays of PhotoSize objects.
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	files.FileReference
	Width  int `json:"width"`
	Height int `json:"height"`
}
