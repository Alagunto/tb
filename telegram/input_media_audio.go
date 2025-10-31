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
