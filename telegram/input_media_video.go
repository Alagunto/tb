package telegram

// InputMediaVideo represents a video to be sent.
//
// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type                  string    `json:"type"`
	Media                 string    `json:"media"`
	Thumbnail             string    `json:"thumbnail,omitempty"`
	Caption               string    `json:"caption,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities       Entities  `json:"caption_entities,omitempty"`
	Width                 int       `json:"width,omitempty"`
	Height                int       `json:"height,omitempty"`
	Duration              int       `json:"duration,omitempty"`
	SupportsStreaming     bool      `json:"supports_streaming,omitempty"`
	HasSpoiler            bool      `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool      `json:"show_caption_above_media,omitempty"`
}

// Field converts InputMediaVideo to a map for API requests.
func (v *InputMediaVideo) Field() (map[string]any, error) {
	result := map[string]any{
		"type":  v.Type,
		"media": v.Media,
	}
	if v.Thumbnail != "" {
		result["thumbnail"] = v.Thumbnail
	}
	if v.Caption != "" {
		result["caption"] = v.Caption
	}
	if v.ParseMode != "" {
		result["parse_mode"] = string(v.ParseMode)
	}
	if len(v.CaptionEntities) > 0 {
		result["caption_entities"] = v.CaptionEntities
	}
	if v.Width > 0 {
		result["width"] = v.Width
	}
	if v.Height > 0 {
		result["height"] = v.Height
	}
	if v.Duration > 0 {
		result["duration"] = v.Duration
	}
	if v.SupportsStreaming {
		result["supports_streaming"] = true
	}
	if v.HasSpoiler {
		result["has_spoiler"] = true
	}
	if v.ShowCaptionAboveMedia {
		result["show_caption_above_media"] = true
	}
	return result, nil
}
