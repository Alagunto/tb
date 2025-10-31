package telegram

// InputMediaPhoto represents a photo to be sent.
//
// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type                  string    `json:"type"`
	Media                 string    `json:"media"`
	Caption               string    `json:"caption,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities       Entities  `json:"caption_entities,omitempty"`
	HasSpoiler            bool      `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool      `json:"show_caption_above_media,omitempty"`
}
