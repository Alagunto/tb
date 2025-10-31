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

// Field converts InputMediaAnimation to a map for API requests.
func (a *InputMediaAnimation) Field() (map[string]any, error) {
	result := map[string]any{
		"type":  a.Type,
		"media": a.Media,
	}
	if a.Thumbnail != "" {
		result["thumbnail"] = a.Thumbnail
	}
	if a.Caption != "" {
		result["caption"] = a.Caption
	}
	if a.ParseMode != "" {
		result["parse_mode"] = string(a.ParseMode)
	}
	if len(a.CaptionEntities) > 0 {
		result["caption_entities"] = a.CaptionEntities
	}
	if a.Width > 0 {
		result["width"] = a.Width
	}
	if a.Height > 0 {
		result["height"] = a.Height
	}
	if a.Duration > 0 {
		result["duration"] = a.Duration
	}
	if a.HasSpoiler {
		result["has_spoiler"] = true
	}
	if a.ShowCaptionAboveMedia {
		result["show_caption_above_media"] = true
	}
	return result, nil
}
