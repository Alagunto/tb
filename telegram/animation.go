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

	// Fields populated when receiving from Telegram (output)
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Duration  int    `json:"duration"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`

	// Optional parameters for sending
	Caption      string `json:"caption,omitempty"`
	HasSpoiler   bool   `json:"has_spoiler,omitempty"`
	CaptionAbove bool   `json:"show_caption_above_media,omitempty"`
}

// Ref returns a FileRef from the received Telegram data.
func (a *Animation) Ref() files.FileRef {
	return files.FileRef{
		FileID:   a.FileID,
		UniqueID: a.UniqueID,
		FileSize: a.FileSize,
	}
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

	b.Add("file_name", a.FileName)

	filesMap := map[string]files.FileSource{
		"animation": a.Source,
	}

	if a.ThumbnailSource != nil {
		filesMap["thumbnail"] = *a.ThumbnailSource
	}

	return &outgoing.Method{
		Name:   "sendAnimation",
		Params: b.Build(),
		Files:  filesMap,
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (a *Animation) UpdateFromResponse(msg *Message) error {
	if msg.Animation != nil {
		a.FileID = msg.Animation.FileID
		a.UniqueID = msg.Animation.UniqueID
		a.Width = msg.Animation.Width
		a.Height = msg.Animation.Height
		a.Duration = msg.Animation.Duration
		a.MIME = msg.Animation.MIME
		a.FileName = msg.Animation.FileName
		a.FileSize = msg.Animation.FileSize
		a.Caption = msg.Caption
	}
	return nil
}
