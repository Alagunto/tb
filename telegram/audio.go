package telegram

import "github.com/alagunto/tb/files"

// Audio represents an audio file to be treated as music by the Telegram clients.
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	files.FileReference
	Duration  int        `json:"duration"`
	Performer string     `json:"performer,omitempty"`
	Title     string     `json:"title,omitempty"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

