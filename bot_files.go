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

// Download saves the file from Telegram servers locally.
// Maximum file size to download is 20 MB.
func (b *Bot[RequestType]) Download(file *files.FileReference, localFilename string) error {
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

// File gets a file from Telegram servers.
func (b *Bot[RequestType]) File(file *files.FileReference) (io.ReadCloser, error) {
	r := NewFileRequester(b.token, b.apiURL, b.client)
	return r.Download(file.FilePath)
}
