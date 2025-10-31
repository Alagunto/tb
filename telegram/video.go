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

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	files.FileReference
	Length    int        `json:"length"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// InputMediaVideo represents a video to be sent.
//
// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type                  string    `json:"type"`
	Media                 string    `json:"media"`
	Thumbnail             string    `json:"thumbnail,omitempty"`
	Caption               string    `json:"caption,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities       Entities  `json:"caption_entities,omitempty"`
	Width                 int       `json:"width,omitempty"`
	Height                int       `json:"height,omitempty"`
	Duration              int       `json:"duration,omitempty"`
	SupportsStreaming     bool      `json:"supports_streaming,omitempty"`
	HasSpoiler            bool      `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool      `json:"show_caption_above_media,omitempty"`
}
