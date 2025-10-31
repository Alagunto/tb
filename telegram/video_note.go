package telegram

import "github.com/alagunto/tb/files"

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	files.FileReference
	Length    int        `json:"length"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}
