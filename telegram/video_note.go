package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// VideoNote object represents a video message.
type VideoNote struct {
	// Source specifies where to get the video note from when sending (input)
	Source files.FileSource `json:"-"`

	// Thumbnail source (optional)
	ThumbnailSource *files.FileSource `json:"-"`

	// Fields populated when receiving from Telegram (output)
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Length    int    `json:"length"`
	Duration  int    `json:"duration"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
}

// Ref returns a FileRef from the received Telegram data.
func (v *VideoNote) Ref() files.FileRef {
	return files.FileRef{
		FileID:   v.FileID,
		UniqueID: v.UniqueID,
		FileSize: v.FileSize,
	}
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (v *VideoNote) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.AddInt("duration", v.Duration)
	b.AddInt("length", v.Length)

	filesMap := map[string]files.FileSource{
		"video_note": v.Source,
	}

	if v.ThumbnailSource != nil {
		filesMap["thumbnail"] = *v.ThumbnailSource
	}

	return &outgoing.Method{
		Name:   "sendVideoNote",
		Params: b.Build(),
		Files:  filesMap,
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (v *VideoNote) UpdateFromResponse(msg *Message) error {
	if msg.VideoNote != nil {
		v.FileID = msg.VideoNote.FileID
		v.UniqueID = msg.VideoNote.UniqueID
		v.Duration = msg.VideoNote.Duration
		v.Length = msg.VideoNote.Length
		v.FileSize = msg.VideoNote.FileSize
	}
	return nil
}
