package tb

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/telegram/methods"
)

// FileByID returns full file object including File.FilePath, allowing you to
// download the file from the server.
//
// Usually, Telegram-provided File objects miss FilePath so you might need to
// perform an additional request to fetch them.
func (b *Bot[RequestType]) FileByID(fileID string) (files.FileReference, error) {
	req := methods.GetFileRequest{
		FileID: fileID,
	}

	r := NewApiRequester[methods.GetFileRequest, methods.GetFileResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getFile", req)
	if err != nil {
		return files.FileReference{}, err
	}
	return *result, nil
}

// Download saves the file from Telegram servers locally.
// Maximum file size to download is 20 MB.
func (b *Bot[RequestType]) Download(file *files.FileRef, localFilename string) error {
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
func (b *Bot[RequestType]) File(file *files.FileRef) (io.ReadCloser, error) {
	var filePath string
	if file.FilePath != "" {
		filePath = file.FilePath
	} else {
		f, err := b.FileByID(file.FileID)
		if err != nil {
			return nil, err
		}
		filePath = f.FilePath
		file.FilePath = filePath // cache the file path
	}

	url := b.apiURL + "/file/bot" + b.token + "/" + filePath

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("telebot: expected status 200 but got %s", resp.Status)
	}

	return resp.Body, nil
}
