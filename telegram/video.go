package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Video object represents a video file.
type Video struct {
	// Source specifies where to get the video from when sending (input)
	Source files.FileSource `json:"-"`

	// Thumbnail source (optional)
	ThumbnailSource *files.FileSource `json:"-"`

	// Fields populated when receiving from Telegram (output)
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Duration  int    `json:"duration"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`

	// Optional parameters for sending
	Caption      string `json:"caption,omitempty"`
	Streaming    bool   `json:"-"`
	HasSpoiler   bool   `json:"has_spoiler,omitempty"`
	CaptionAbove bool   `json:"show_caption_above_media,omitempty"`
}

// Ref returns a FileRef from the received Telegram data.
func (v *Video) Ref() files.FileRef {
	return files.FileRef{
		FileID:   v.FileID,
		UniqueID: v.UniqueID,
		FileSize: v.FileSize,
	}
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (v *Video) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()

	mediaParams := &params.MediaParams{
		Caption:      v.Caption,
		HasSpoiler:   v.HasSpoiler,
		CaptionAbove: v.CaptionAbove,
	}
	mediaParams.Apply(b)

	dimParams := &params.DimensionParams{
		Width:    v.Width,
		Height:   v.Height,
		Duration: v.Duration,
	}
	dimParams.Apply(b)

	b.AddBool("supports_streaming", v.Streaming)
	b.Add("file_name", v.FileName)

	filesMap := map[string]files.FileSource{
		"video": v.Source,
	}

	if v.ThumbnailSource != nil {
		filesMap["thumbnail"] = *v.ThumbnailSource
	}

	return &outgoing.Method{
		Name:   "sendVideo",
		Params: b.Build(),
		Files:  filesMap,
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (v *Video) UpdateFromResponse(msg interface{}) error {
	m, ok := msg.(*Message)
	if !ok || m.Video == nil {
		return nil
	}
	v.FileID = m.Video.FileID
	v.UniqueID = m.Video.UniqueID
	v.Width = m.Video.Width
	v.Height = m.Video.Height
	v.Duration = m.Video.Duration
	v.MIME = m.Video.MIME
	v.FileName = m.Video.FileName
	v.FileSize = m.Video.FileSize
	v.Caption = m.Caption
	return nil
}
