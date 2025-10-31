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

// InputMediaPhoto represents a photo to be sent.
//
// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type                  string    `json:"type"`
	Media                 string    `json:"media"`
	Caption               string    `json:"caption,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities       Entities  `json:"caption_entities,omitempty"`
	HasSpoiler            bool      `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool      `json:"show_caption_above_media,omitempty"`
}
