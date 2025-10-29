package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Sticker object represents a sticker.
type Sticker struct {
	// Source specifies where to get the sticker from when sending (input)
	Source files.FileSource `json:"-"`

	// Fields populated when receiving from Telegram (output)
	FileID        string        `json:"file_id"`
	UniqueID      string        `json:"file_unique_id"`
	Type          string        `json:"type"`
	Width         int           `json:"width"`
	Height        int           `json:"height"`
	Animated      bool          `json:"is_animated"`
	Video         bool          `json:"is_video"`
	Thumbnail     *Photo        `json:"thumbnail,omitempty"`
	Emoji         string        `json:"emoji,omitempty"`
	SetName       string        `json:"set_name,omitempty"`
	MaskPosition  *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID string        `json:"custom_emoji_id,omitempty"`
	Repaint       bool          `json:"needs_repainting,omitempty"`
	FileSize      int64         `json:"file_size,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// Ref returns a FileRef from the received Telegram data.
func (s *Sticker) Ref() files.FileRef {
	return files.FileRef{
		FileID:   s.FileID,
		UniqueID: s.UniqueID,
		FileSize: s.FileSize,
	}
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (s *Sticker) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.Add("emoji", s.Emoji)

	return &outgoing.Method{
		Name:   "sendSticker",
		Params: b.Build(),
		Files: map[string]files.FileSource{
			"sticker": s.Source,
		},
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (s *Sticker) UpdateFromResponse(msg *Message) error {
	if msg.Sticker != nil {
		s.FileID = msg.Sticker.FileID
		s.UniqueID = msg.Sticker.UniqueID
		s.Type = msg.Sticker.Type
		s.Width = msg.Sticker.Width
		s.Height = msg.Sticker.Height
		s.Animated = msg.Sticker.Animated
		s.Video = msg.Sticker.Video
		s.Emoji = msg.Sticker.Emoji
		s.SetName = msg.Sticker.SetName
		s.CustomEmojiID = msg.Sticker.CustomEmojiID
		s.Repaint = msg.Sticker.Repaint
		s.FileSize = msg.Sticker.FileSize
	}
	return nil
}
