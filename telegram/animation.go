package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Animation object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// Source specifies where to get the animation from when sending (input)
	Source files.FileSource `json:"-"`

	// Thumbnail source (optional)
	ThumbnailSource *files.FileSource `json:"-"`

	Thumbnail *Photo `json:"thumbnail,omitempty"`

	// Optional parameters for sending
	Caption      string `json:"caption,omitempty"`
	HasSpoiler   bool   `json:"has_spoiler,omitempty"`
	CaptionAbove bool   `json:"show_caption_above_media,omitempty"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Duration     int    `json:"duration"`

	// Telegram-injected data
	ref files.FileRef
}

// Ref returns a FileRef from the received Telegram data.
func (a *Animation) Ref() files.FileRef {
	return a.ref
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (a *Animation) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()

	mediaParams := &params.MediaParams{
		Caption:      a.Caption,
		HasSpoiler:   a.HasSpoiler,
		CaptionAbove: a.CaptionAbove,
	}
	mediaParams.Apply(b)

	dimParams := &params.DimensionParams{
		Width:    a.Width,
		Height:   a.Height,
		Duration: a.Duration,
	}
	dimParams.Apply(b)

	b.Add("media", a.Source.GetFilenameForUpload())

	files := map[string]files.FileSource{
		a.Source.GetFilenameForUpload(): a.Source,
	}
	if a.ThumbnailSource != nil {
		files[a.ThumbnailSource.GetFilenameForUpload()] = *a.ThumbnailSource
	}

	return &outgoing.Method{
		Name:   "sendAnimation",
		Params: b.Build(),
		Files:  files,
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (a *Animation) UpdateFromResponse(msg *Message) error {
	if msg.Animation != nil {
		a.ref = msg.Animation.Ref()
	}
	return nil
}
