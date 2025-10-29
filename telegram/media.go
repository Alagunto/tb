package telegram

// Photo object represents a single photo file.
type Photo struct {
	FileID   string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int64  `json:"file_size,omitempty"`
}

// Audio object represents an audio file.
type Audio struct {
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
}

// Document object represents a general file.
type Document struct {
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
}

// Sticker object represents a sticker.
type Sticker struct {
	FileID        string        `json:"file_id"`
	UniqueID      string        `json:"file_unique_id"`
	Type          string        `json:"type"`
	Width         int           `json:"width"`
	Height        int           `json:"height"`
	Animated      bool          `json:"is_animated"`
	Video         bool          `json:"is_video"`
	Thumbnail     *Photo        `json:"thumbnail,omitempty"`
	Emoji         string        `json:"emoji,omitempty"`
	SetName       string        `json:"set_name,omitempty"`
	MaskPosition  *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID string        `json:"custom_emoji_id,omitempty"`
	Repaint       bool          `json:"needs_repainting,omitempty"`
	FileSize      int64         `json:"file_size,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// Video object represents a video file.
type Video struct {
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Duration  int    `json:"duration"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
}

// Animation object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Duration  int    `json:"duration"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
}

// Voice object represents a voice note.
type Voice struct {
	FileID   string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	Duration int    `json:"duration"`
	MIME     string `json:"mime_type,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}

// VideoNote object represents a video message.
type VideoNote struct {
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Length    int    `json:"length"`
	Duration  int    `json:"duration"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

// Venue represents a venue.
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    string   `json:"foursquare_id,omitempty"`
	FoursquareType  string   `json:"foursquare_type,omitempty"`
	GooglePlaceID   string   `json:"google_place_id,omitempty"`
	GooglePlaceType string   `json:"google_place_type,omitempty"`
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// Story represents a story published by a user.
type Story struct {
	Chat *Chat `json:"chat"`
	ID   int   `json:"story_id"`
}
