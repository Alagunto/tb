package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SourceType represents the type of file source.
type SourceType int

const (
	// SourceUnknown indicates an uninitialized or invalid source
	SourceUnknown SourceType = iota
	// SourceTelegramFile indicates a file already on Telegram servers (file_id)
	SourceTelegramFile
	// SourceURL indicates a file accessible via HTTP/HTTPS URL
	SourceURL
	// SourceLocalFile indicates a file on the local filesystem
	SourceLocalFile
	// SourceReader indicates a file from an io.Reader
	SourceReader
)

// FileSource represents where a file comes from for upload to Telegram.
// Only one of the fields should be set at a time.
type FileSource struct {
	// TelegramFileID is the file_id for files already on Telegram servers
	TelegramFileID string

	// URL is the HTTP/HTTPS URL for remote files
	URL string

	// LocalPath is the filesystem path for local files
	LocalPath string

	// Reader is an io.Reader for streaming file content
	Reader io.Reader

	// Filename is required when using Reader
	Filename string
}

// UseTelegramFile creates a FileSource from a Telegram file_id.
func UseTelegramFile(fileID string) FileSource {
	return FileSource{TelegramFileID: fileID}
}

// UseURL creates a FileSource from a URL.
func UseURL(url string) FileSource {
	return FileSource{URL: url}
}

// UseLocalFile creates a FileSource from a local file path.
func UseLocalFile(path string) FileSource {
	return FileSource{LocalPath: path}
}

// UseReader creates a FileSource from an io.Reader with the given filename.
func UseReader(reader io.Reader, filename string) FileSource {
	return FileSource{
		Reader:   reader,
		Filename: filename,
	}
}

// Type returns the SourceType based on which field is set.
func (fs *FileSource) Type() SourceType {
	if fs.TelegramFileID != "" {
		return SourceTelegramFile
	}
	if fs.URL != "" {
		return SourceURL
	}
	if fs.LocalPath != "" {
		return SourceLocalFile
	}
	if fs.Reader != nil {
		return SourceReader
	}
	return SourceUnknown
}

// Validate checks if the FileSource is properly configured.
func (fs *FileSource) Validate() error {
	switch fs.Type() {
	case SourceUnknown:
		return fmt.Errorf("file source not initialized")
	case SourceTelegramFile:
		if fs.TelegramFileID == "" {
			return fmt.Errorf("telegram file_id is empty")
		}
	case SourceURL:
		if fs.URL == "" {
			return fmt.Errorf("URL is empty")
		}
	case SourceLocalFile:
		if fs.LocalPath == "" {
			return fmt.Errorf("local path is empty")
		}
		// Check if file exists
		if _, err := os.Stat(fs.LocalPath); err != nil {
			return fmt.Errorf("local file does not exist: %w", err)
		}
	case SourceReader:
		if fs.Reader == nil {
			return fmt.Errorf("reader is nil")
		}
		if fs.Filename == "" {
			return fmt.Errorf("filename is required when using reader")
		}
	}
	return nil
}

// ToTelegramParam converts the FileSource to a Telegram API parameter value.
// Returns the parameter value, whether upload is needed, and an error.
func (fs *FileSource) ToTelegramParam(filename string) (paramValue string, needsUpload bool, err error) {
	if err := fs.Validate(); err != nil {
		return "", false, err
	}

	switch fs.Type() {
	case SourceTelegramFile:
		return fs.TelegramFileID, false, nil
	case SourceURL:
		return fs.URL, false, nil
	case SourceLocalFile:
		// For local files, we use attach:// protocol
		return fmt.Sprintf("attach://%s", filename), true, nil
	case SourceReader:
		// For readers, we also use attach:// protocol
		return fmt.Sprintf("attach://%s", filename), true, nil
	default:
		return "", false, fmt.Errorf("unknown source type")
	}
}

// GetFilenameForUpload returns the filename to use when uploading.
func (fs *FileSource) GetFilenameForUpload() string {
	switch fs.Type() {
	case SourceLocalFile:
		return filepath.Base(fs.LocalPath)
	case SourceReader:
		return fs.Filename
	default:
		return ""
	}
}

// FileRef represents a file reference returned by Telegram after upload.
type FileRef struct {
	// FileID is the unique identifier for the file on Telegram servers
	FileID string `json:"file_id"`

	// UniqueID is the unique identifier for the file, which is supposed to be
	// the same over time and for different bots
	UniqueID string `json:"file_unique_id"`

	// FileSize is the file size in bytes (if known)
	FileSize int64 `json:"file_size,omitempty"`

	// FilePath is the file path on Telegram servers (used for downloading)
	FilePath string `json:"file_path,omitempty"`

	// FileURL is the HTTP/HTTPS URL for remote files
	FileURL string

	// FileLocal is the local filesystem path for downloaded files
	FileLocal string

	// FileReader is an io.Reader for streaming file content
	FileReader io.Reader
}

// IsEmpty checks if the FileRef is uninitialized.
func (fr *FileRef) IsEmpty() bool {
	return fr.FileID == ""
}

// AsSource converts the FileRef back to a FileSource for re-sending.
func (fr *FileRef) AsSource() FileSource {
	return UseTelegramFile(fr.FileID)
}

// InCloud returns true if the file exists on Telegram servers (has FileID)
func (fr *FileRef) InCloud() bool {
	return fr.FileID != ""
}

// OnDisk returns true if the file exists on the local filesystem
func (fr *FileRef) OnDisk() bool {
	return fr.FileLocal != ""
}
