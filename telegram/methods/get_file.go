package methods

import "github.com/alagunto/tb/files"

// GetFileRequest represents the request for getFile method.
type GetFileRequest struct {
	// File identifier to get information about
	FileID string `json:"file_id"`
}

// GetFileResponse represents the response for getFile method.
type GetFileResponse = files.FileRef
