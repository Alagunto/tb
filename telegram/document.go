package telegram

import (
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Document object represents a general file.
type Document struct {
	// Source specifies where to get the document from when sending (input)
	Source files.FileSource `json:"-"`

	// Thumbnail source (optional)
	ThumbnailSource *files.FileSource `json:"-"`

	// Fields populated when receiving from Telegram (output)
	FileID    string `json:"file_id"`
	UniqueID  string `json:"file_unique_id"`
	Thumbnail *Photo `json:"thumbnail,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MIME      string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`

	// Optional parameters for sending
	Caption              string `json:"caption,omitempty"`
	DisableTypeDetection bool   `json:"-"`
}

// Ref returns a FileRef from the received Telegram data.
func (d *Document) Ref() files.FileRef {
	return files.FileRef{
		FileID:   d.FileID,
		UniqueID: d.UniqueID,
		FileSize: d.FileSize,
	}
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (d *Document) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.Add("caption", d.Caption)
	b.Add("file_name", d.FileName)
	b.AddBool("disable_content_type_detection", d.DisableTypeDetection)

	filesMap := map[string]files.FileSource{
		"document": d.Source,
	}

	if d.ThumbnailSource != nil {
		filesMap["thumbnail"] = *d.ThumbnailSource
	}

	return &outgoing.Method{
		Name:   "sendDocument",
		Params: b.Build(),
		Files:  filesMap,
	}
}

// UpdateFromResponse implements the outgoing.ResponseHandler interface.
func (d *Document) UpdateFromResponse(msg *Message) error {
	if msg.Document != nil {
		d.FileID = msg.Document.FileID
		d.UniqueID = msg.Document.UniqueID
		d.MIME = msg.Document.MIME
		d.FileName = msg.Document.FileName
		d.FileSize = msg.Document.FileSize
		d.Caption = msg.Caption
	}
	return nil
}
