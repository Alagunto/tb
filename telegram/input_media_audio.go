package telegram

// InputMediaAudio represents an audio file to be treated as music to be sent.
//
// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            string    `json:"type"`
	Media           string    `json:"media"`
	Thumbnail       string    `json:"thumbnail,omitempty"`
	Caption         string    `json:"caption,omitempty"`
	ParseMode       ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities Entities  `json:"caption_entities,omitempty"`
	Duration        int       `json:"duration,omitempty"`
	Performer       string    `json:"performer,omitempty"`
	Title           string    `json:"title,omitempty"`
}

// Field converts InputMediaAudio to a map for API requests.
func (a *InputMediaAudio) Field() (map[string]any, error) {
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
	if a.Duration > 0 {
		result["duration"] = a.Duration
	}
	if a.Performer != "" {
		result["performer"] = a.Performer
	}
	if a.Title != "" {
		result["title"] = a.Title
	}
	return result, nil
}
