package telegram

import "github.com/alagunto/tb/files"

// Voice represents a voice note.
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	files.FileReference
	Duration int `json:"duration"`
}

