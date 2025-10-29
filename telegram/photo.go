package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Photo object represents a single photo file.
// It can be used both for sending (with Source) and receiving (FileID/UniqueID populated).
type Photo struct {
	// Source specifies where to get the photo from when sending (input)
	Source files.FileSource `json:"-"`

	// Fields populated when receiving from Telegram (output)
	FileID   string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	FileSize int64  `json:"file_size,omitempty"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`

	// Optional parameters for sending
	Caption      string `json:"caption,omitempty"`
	HasSpoiler   bool   `json:"has_spoiler,omitempty"`
	CaptionAbove bool   `json:"show_caption_above_media,omitempty"`
}

// Ref returns a FileRef from the received Telegram data.
func (p *Photo) Ref() files.FileRef {
	return files.FileRef{
		FileID:   p.FileID,
		UniqueID: p.UniqueID,
		FileSize: p.FileSize,
	}
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (p *Photo) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()

	mediaParams := &params.MediaParams{
		Caption:      p.Caption,
		HasSpoiler:   p.HasSpoiler,
		CaptionAbove: p.CaptionAbove,
	}
	mediaParams.Apply(b)

	return &outgoing.Method{
		Name:   "sendPhoto",
		Params: b.Build(),
		Files: map[string]files.FileSource{
			"photo": p.Source,
		},
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (p *Photo) UpdateFromResponse(msg interface{}) error {
	m, ok := msg.(*Message)
	if !ok || m.Photo == nil {
		return nil
	}
	p.FileID = m.Photo.FileID
	p.UniqueID = m.Photo.UniqueID
	p.FileSize = m.Photo.FileSize
	p.Width = m.Photo.Width
	p.Height = m.Photo.Height
	p.Caption = m.Caption
	return nil
}
