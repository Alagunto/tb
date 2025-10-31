package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

type InputAlbum struct {
	Media        []outgoing.Content `json:"media"`
	Caption      string             `json:"caption,omitempty"`
	ParseMode    ParseMode          `json:"parse_mode,omitempty"`
	Entities     []MessageEntity    `json:"caption_entities,omitempty"`
	CaptionAbove bool               `json:"show_caption_above_media,omitempty"`
	HasSpoiler   bool               `json:"has_spoiler,omitempty"`
	StarCount    int                `json:"star_count,omitempty"`
}

func (i *InputAlbum) ToTelegramSendMethod() *outgoing.Method {
	mediaElements := make([]any, len(i.Media))
	for _, media := range i.Media {
		mediaElements = append(mediaElements, media.ToTelegramSendMethod().Params)
	}

	files := make(map[string]files.FileSource)
	for _, media := range i.Media {
		for name, source := range media.ToTelegramSendMethod().Files {
			files[name] = source
		}
	}

	return &outgoing.Method{
		Name: "sendMediaGroup",
		Params: params.New().
			Add("media", mediaElements).
			Build(),
		Files: files,
	}
}
