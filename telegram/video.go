package telegram

import "github.com/alagunto/tb/files"

// Video represents a video file.
//
// https://core.telegram.org/bots/api#video
type Video struct {
	files.FileReference
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}
