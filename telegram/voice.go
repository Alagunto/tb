package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Voice object represents a voice note.
type Voice struct {
	// Source specifies where to get the voice from when sending (input)
	Source files.FileSource `json:"-"`

	// Fields populated when receiving from Telegram (output)
	FileID   string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	Duration int    `json:"duration"`
	MIME     string `json:"mime_type,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`

	// Optional parameters for sending
	Caption string `json:"caption,omitempty"`
}

// Ref returns a FileRef from the received Telegram data.
func (v *Voice) Ref() files.FileRef {
	return files.FileRef{
		FileID:   v.FileID,
		UniqueID: v.UniqueID,
		FileSize: v.FileSize,
	}
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (v *Voice) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.Add("caption", v.Caption)
	b.AddInt("duration", v.Duration)

	return &outgoing.Method{
		Name:   "sendVoice",
		Params: b.Build(),
		Files: map[string]files.FileSource{
			"voice": v.Source,
		},
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (v *Voice) UpdateFromResponse(msg interface{}) error {
	m, ok := msg.(*Message)
	if !ok || m.Voice == nil {
		return nil
	}
	v.FileID = m.Voice.FileID
	v.UniqueID = m.Voice.UniqueID
	v.Duration = m.Voice.Duration
	v.MIME = m.Voice.MIME
	v.FileSize = m.Voice.FileSize
	v.Caption = m.Caption
	return nil
}
