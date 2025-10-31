package files

import "io"

// FileReference represents a file reference returned by Telegram after upload.
type FileReference struct {
	// FileID is the unique identifier for the file on Telegram servers
	ID string `json:"file_id"`

	// UniqueID is the unique identifier for the file, which is supposed to be
	// the same over time and for different bots
	UniqueID string `json:"file_unique_id"`

	// FileSize is the file size in bytes (if known)
	FileSize int64 `json:"file_size,omitempty"`

	// FilePath is the file path on Telegram servers (used for downloading)
	FilePath string `json:"file_path,omitempty"`

	// MIME is the MIME type of the file
	MIME string `json:"mime_type,omitempty"`

	// FileName is the name of the file
	FileName string `json:"file_name,omitempty"`

	// FileURL is the HTTP/HTTPS URL for remote files
	FileURL string

	// FileLocal is the local filesystem path for downloaded files
	FileLocal string

	// FileReader is an io.Reader for streaming file content
	FileReader io.Reader
}

// IsEmpty checks if the FileReference is uninitialized.
func (fr *FileReference) IsEmpty() bool {
	return fr.ID == ""
}

// AsSource converts the FileReference back to a FileSource for re-sending.
func (fr *FileReference) AsSource() FileSource {
	return UseTelegramFile(fr.ID)
}

// InCloud returns true if the file exists on Telegram servers (has FileID)
func (fr *FileReference) InCloud() bool {
	return fr.ID != ""
}

// OnDisk returns true if the file exists on the local filesystem
func (fr *FileReference) OnDisk() bool {
	return fr.FileLocal != ""
}
