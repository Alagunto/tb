package tb

import (
	"context"
	"io"
	"os"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/telegram"
)

// FileByID returns full file object including File.FilePath, allowing you to
// download the file from the server.
//
// Usually, Telegram-provided File objects miss FilePath so you might need to
// perform an additional request to fetch them.
func (b *Bot[RequestType]) FileByID(fileID string) (files.FileReference, error) {
	if fileID == "" {
		return files.FileReference{}, errors.WithInvalidParam(errors.ErrTelebot, "file_id", nil)
	}

	req := telegram.GetFileRequest{
		FileID: fileID,
	}

	r := NewApiRequester[telegram.GetFileRequest, files.FileReference](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getFile", req)
	if err != nil {
		return files.FileReference{}, err
	}
	return *result, nil
}

// FileByIDBackground is a convenience wrapper that uses context.Background()
func (b *Bot[RequestType]) FileByIDBackground(fileID string) (files.FileReference, error) {
	return b.FileByID(fileID)
}

// Download saves the file from Telegram servers locally.
// Maximum file size to download is 20 MB.
func (b *Bot[RequestType]) Download(file *files.FileReference, localFilename string) error {
	if file == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "file", nil)
	}
	if localFilename == "" {
		return errors.WithInvalidParam(errors.ErrTelebot, "local_filename", nil)
	}

	reader, err := b.File(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	out, err := os.Create(localFilename)
	if err != nil {
		return errors.Wrap(err)
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// DownloadBackground is a convenience wrapper that uses context.Background()
func (b *Bot[RequestType]) DownloadBackground(file *files.FileReference, localFilename string) error {
	return b.Download(file, localFilename)
}

// File gets a file from Telegram servers.
func (b *Bot[RequestType]) File(file *files.FileReference) (io.ReadCloser, error) {
	if file == nil {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "file", nil)
	}

	r := NewFileRequester(b.token, b.apiURL, b.client)
	return r.Download(file.FilePath)
}

// FileBackground is a convenience wrapper that uses context.Background()
func (b *Bot[RequestType]) FileBackground(file *files.FileReference) (io.ReadCloser, error) {
	return b.File(file)
}
