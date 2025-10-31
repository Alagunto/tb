package telegram

import "github.com/alagunto/tb/files"

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
//
// https://core.telegram.org/bots/api#animation
type Animation struct {
	files.FileReference
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}
