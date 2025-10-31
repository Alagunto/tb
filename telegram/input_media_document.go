package telegram

// InputMediaDocument represents a general file to be sent.
//
// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        string    `json:"type"`
	Media                       string    `json:"media"`
	Thumbnail                   string    `json:"thumbnail,omitempty"`
	Caption                     string    `json:"caption,omitempty"`
	ParseMode                   ParseMode `json:"parse_mode,omitempty"`
	CaptionEntities             Entities  `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool      `json:"disable_content_type_detection,omitempty"`
}

// Field converts InputMediaDocument to a map for API requests.
func (d *InputMediaDocument) Field() (map[string]any, error) {
	result := map[string]any{
		"type":  d.Type,
		"media": d.Media,
	}
	if d.Thumbnail != "" {
		result["thumbnail"] = d.Thumbnail
	}
	if d.Caption != "" {
		result["caption"] = d.Caption
	}
	if d.ParseMode != "" {
		result["parse_mode"] = string(d.ParseMode)
	}
	if len(d.CaptionEntities) > 0 {
		result["caption_entities"] = d.CaptionEntities
	}
	if d.DisableContentTypeDetection {
		result["disable_content_type_detection"] = true
	}
	return result, nil
}
