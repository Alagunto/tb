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
