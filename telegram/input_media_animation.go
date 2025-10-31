package telegram

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
//
// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type                  string    `json:"type"`
	Media                 string    `json:"media"`
	Thumbnail             string    `json:"thumbnail,omitempty"`
	Caption               string    `json:"caption,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities       Entities  `json:"caption_entities,omitempty"`
	Width                 int       `json:"width,omitempty"`
	Height                int       `json:"height,omitempty"`
	Duration              int       `json:"duration,omitempty"`
	HasSpoiler            bool      `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool      `json:"show_caption_above_media,omitempty"`
}
