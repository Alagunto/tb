package media

import "github.com/alagunto/tb/files"

// Sticker represents a sticker.
//
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	files.FileReference
	Type            string     `json:"type"`
	Width           int        `json:"width"`
	Height          int        `json:"height"`
	IsAnimated      bool       `json:"is_animated"`
	IsVideo         bool       `json:"is_video"`
	Thumbnail       *PhotoSize `json:"thumbnail,omitempty"`
	Emoji           string     `json:"emoji,omitempty"`
	SetName         string     `json:"set_name,omitempty"`
	PremiumAnimation *Animation `json:"premium_animation,omitempty"`
	MaskPosition    *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID   string     `json:"custom_emoji_id,omitempty"`
	NeedsRepainting bool       `json:"needs_repainting,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
//
// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

