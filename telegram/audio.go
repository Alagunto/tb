package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Audio object represents an audio file.
type Audio struct {
	// Source specifies where to get the audio from when sending (input)
	Source files.FileSource `json:"-"`

	// Thumbnail source (optional)
	ThumbnailSource *files.FileSource `json:"-"`

	// Fields populated when receiving from Telegram (output)
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`

	// Optional parameters for sending
	Caption string `json:"caption,omitempty"`
}

// Ref returns a FileRef from the received Telegram data.
func (a *Audio) Ref() files.FileRef {
	return files.FileRef{
		FileID:   a.FileID,
		UniqueID: a.UniqueID,
		FileSize: a.FileSize,
	}
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (a *Audio) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.Add("caption", a.Caption)
	b.AddInt("duration", a.Duration)
	b.Add("title", a.Title)
	b.Add("performer", a.Performer)
	b.Add("file_name", a.FileName)

	filesMap := map[string]files.FileSource{
		"audio": a.Source,
	}

	if a.ThumbnailSource != nil {
		filesMap["thumbnail"] = *a.ThumbnailSource
	}

	return &outgoing.Method{
		Name:   "sendAudio",
		Params: b.Build(),
		Files:  filesMap,
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (a *Audio) UpdateFromResponse(msg interface{}) error {
	m, ok := msg.(*Message)
	if !ok || m.Audio == nil {
		return nil
	}
	a.FileID = m.Audio.FileID
	a.UniqueID = m.Audio.UniqueID
	a.Duration = m.Audio.Duration
	a.Title = m.Audio.Title
	a.Performer = m.Audio.Performer
	a.MIME = m.Audio.MIME
	a.FileName = m.Audio.FileName
	a.FileSize = m.Audio.FileSize
	a.Caption = m.Caption
	return nil
}
