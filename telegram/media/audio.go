package media

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/telegram"
)

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

// InputMediaAudio represents an audio file to be treated as music to be sent.
//
// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            string                `json:"type"`
	Media           string                `json:"media"`
	Thumbnail       string                `json:"thumbnail,omitempty"`
	Caption         string                `json:"caption,omitempty"`
	ParseMode       telegram.ParseMode     `json:"parse_mode,omitempty"`
	CaptionEntities telegram.Entities      `json:"caption_entities,omitempty"`
	Duration        int                    `json:"duration,omitempty"`
	Performer       string                `json:"performer,omitempty"`
	Title           string                `json:"title,omitempty"`
}

