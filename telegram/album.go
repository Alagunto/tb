package telegram

// Media represents a media file that can be sent to Telegram.
type Media interface {
	MediaType() string
}

// Inputtable is an interface for media types that can be sent or edited.
type Inputtable interface {
	InputMedia() *InputMedia
	MediaFile() interface{}
}

// Album is a slice of media items (photos or videos) that can be sent as a group.
type Album []Inputtable

// PaidAlbum is a slice of paid media items that can be sent as a group.
type PaidAlbum []PaidInputtable

// PaidInputtable is an interface for paid media types.
type PaidInputtable interface {
	InputMedia() *InputMedia
	MediaFile() interface{}
}

// InputMedia represents the content of a media message to be sent.
type InputMedia struct {
	Type  string `json:"type"`
	Media string `json:"media"`

	Caption              string          `json:"caption,omitempty"`
	ParseMode            ParseMode       `json:"parse_mode,omitempty"`
	Entities             []MessageEntity `json:"caption_entities,omitempty"`
	CaptionAbove         bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler           bool            `json:"has_spoiler,omitempty"`
	Thumbnail            string          `json:"thumbnail,omitempty"`
	Width                int             `json:"width,omitempty"`
	Height               int             `json:"height,omitempty"`
	Duration             int             `json:"duration,omitempty"`
	SupportsStreaming    bool            `json:"supports_streaming,omitempty"`
	Performer            string          `json:"performer,omitempty"`
	Title                string          `json:"title,omitempty"`
	DisableTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}
