package telegram

import "github.com/alagunto/tb/files"

// Document represents a general file (as opposed to photos, voice messages and audio files).
//
// https://core.telegram.org/bots/api#document
type Document struct {
	files.FileReference
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}
