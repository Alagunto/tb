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

// Field converts InputMediaPhoto to a map for API requests.
func (p *InputMediaPhoto) Field() (map[string]any, error) {
	result := map[string]any{
		"type":  p.Type,
		"media": p.Media,
	}
	if p.Caption != "" {
		result["caption"] = p.Caption
	}
	if p.ParseMode != "" {
		result["parse_mode"] = string(p.ParseMode)
	}
	if len(p.CaptionEntities) > 0 {
		result["caption_entities"] = p.CaptionEntities
	}
	if p.HasSpoiler {
		result["has_spoiler"] = true
	}
	if p.ShowCaptionAboveMedia {
		result["show_caption_above_media"] = true
	}
	return result, nil
}
