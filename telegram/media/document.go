package media

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/telegram"
)

// Document represents a general file (as opposed to photos, voice messages and audio files).
//
// https://core.telegram.org/bots/api#document
type Document struct {
	files.FileReference
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// InputMediaDocument represents a general file to be sent.
//
// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        string             `json:"type"`
	Media                       string             `json:"media"`
	Thumbnail                   string             `json:"thumbnail,omitempty"`
	Caption                     string             `json:"caption,omitempty"`
	ParseMode                   telegram.ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities             telegram.Entities  `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool               `json:"disable_content_type_detection,omitempty"`
}
