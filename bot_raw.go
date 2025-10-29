package tb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/telegram"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Raw lets you call any method of Bot API manually.
// It also handles API errors, so you only need to unwrap
// result field from json data.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Raw(method string, payload interface{}) ([]byte, error) {
	// Apply rate limiting if configured
	if b.rateLimiter != nil {
		b.rateLimiter.Wait()
	}

	// Define the actual API call
	apiCall := func() ([]byte, error) {
		return b.rawAPICall(method, payload)
	}

	// Apply retry policy if configured
	if b.retryPolicy != nil {
		return WithRetry(apiCall, *b.retryPolicy)
	}

	return apiCall()
}

// rawAPICall performs the actual HTTP request to Telegram API.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) rawAPICall(method string, payload interface{}) ([]byte, error) {
	url := b.apiURL + "/bot" + b.token + "/" + method

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return nil, err
	}

	// Cancel the request immediately without waiting for the timeout
	// when bot is about to stop.
	// This may become important if doing long polling with long timeout.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		b.stopMu.RLock()
		stopCh := b.stopClient
		b.stopMu.RUnlock()

		select {
		case <-stopCh:
			cancel()
		case <-ctx.Done():
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	resp.Close = true
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if b.settings.Verbose {
		verbose(method, payload, data)
	}

	// returning data as well
	return data, extractOk(data)
}

func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) RawSendFiles(method string, filesToSend map[string]files.FileSource, params map[string]string) ([]byte, error) {
	// Apply rate limiting if configured
	if b.rateLimiter != nil {
		b.rateLimiter.Wait()
	}

	rawFiles := make(map[string]interface{})
	fileNames := make(map[string]string)

	for name, source := range filesToSend {
		paramValue, needsUpload, err := source.ToTelegramParam(name)
		if err != nil {
			return nil, errors.Wrap(err)
		}

		if needsUpload {
			// File needs to be uploaded
			switch source.Type() {
			case files.SourceLocalFile:
				rawFiles[name] = source.LocalPath
			case files.SourceTelegramFile:
				rawFiles[name] = source.Reader
			default:
				return nil, fmt.Errorf("telebot: unsupported file source type for field %s", name)
			}
			fileNames[name] = source.GetFilenameForUpload()
		} else {
			// File is already on Telegram or accessible via URL
			params[name] = paramValue
		}
	}

	if len(rawFiles) == 0 {
		return b.Raw(method, params)
	}

	// File upload logic - wrap in retry if configured
	uploadFunc := func() ([]byte, error) {
		return b.sendFilesWithMultipart(method, rawFiles, fileNames, params)
	}

	if b.retryPolicy != nil {
		return WithRetry(uploadFunc, *b.retryPolicy)
	}

	return uploadFunc()
}

// sendFilesWithMultipart performs the actual file upload via multipart/form-data
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) sendFilesWithMultipart(method string, rawFiles map[string]interface{}, fileNames map[string]string, params map[string]string) ([]byte, error) {
	pipeReader, pipeWriter := io.Pipe()
	writer := multipart.NewWriter(pipeWriter)

	go func() {
		defer pipeWriter.Close()

		for field, file := range rawFiles {
			if err := addFileToWriter(writer, fileNames[field], field, file); err != nil {
				pipeWriter.CloseWithError(err)
				return
			}
		}
		for field, value := range params {
			if err := writer.WriteField(field, value); err != nil {
				pipeWriter.CloseWithError(err)
				return
			}
		}
		if err := writer.Close(); err != nil {
			pipeWriter.CloseWithError(err)
			return
		}
	}()

	url := b.apiURL + "/bot" + b.token + "/" + method

	resp, err := b.client.Post(url, writer.FormDataContentType(), pipeReader)
	if err != nil {
		err = errors.Wrap(err)
		pipeReader.CloseWithError(err)
		return nil, err
	}
	resp.Close = true
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.ErrTelegramInternal
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return data, extractOk(data)
}

func addFileToWriter(writer *multipart.Writer, filename, field string, file interface{}) error {
	var reader io.Reader
	if r, ok := file.(io.Reader); ok {
		reader = r
	} else if path, ok := file.(string); ok {
		f, err := os.Open(path)
		if err != nil {
			return errors.Wrap(err)
		}
		defer f.Close()
		reader = f
	} else {
		// TODO: fix error
		return errors.WithInvalidParam(errors.ErrUnsupportedWhat, "file", fmt.Sprintf("%v", file))
	}

	part, err := writer.CreateFormFile(field, filename)
	if err != nil {
		return errors.Wrap(err)
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) RawSendText(to bot.Recipient, text string, opt communications.SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id": to.Recipient(),
		"text":    text,
	}
	b.RawEmbedSendOptions(params, opt)

	data, err := b.Raw("sendMessage", params)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var resp struct {
		Result telegram.Message `json:"result"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err)
	}
	return &resp.Result, nil
}

func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) RawSendMedia(media Media, params map[string]string, filesToSend map[string]files.FileSource) (*Message, error) {
	kind := media.MediaType()
	caser := cases.Title(language.English)
	what := "send" + caser.String(kind)

	if kind == "videoNote" {
		kind = "video_note"
	}

	data, err := b.RawSendFiles(what, filesToSend, params)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var resp struct {
		Result telegram.Message `json:"result"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err)
	}
	return &resp.Result, nil
}

func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) getMe() (*telegram.User, error) {
	data, err := b.Raw("getMe", nil)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var resp struct {
		Result *telegram.User
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err)
	}
	return resp.Result, nil
}

// GetUpdates fetches updates from the Telegram API.
// Do not use this method directly by default, instead use Start() to start the Poller to fetch updates automatically.
// Use it only if you need to fetch updates manually, without starting the bot as usual.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) GetUpdates(offset, limit int, timeout time.Duration, allowed []string) ([]telegram.Update, error) {
	params := map[string]string{
		"offset":  strconv.Itoa(offset),
		"timeout": strconv.Itoa(int(timeout / time.Second)),
	}

	data, _ := json.Marshal(allowed)
	params["allowed_updates"] = string(data)

	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	data, err := b.Raw("getUpdates", params)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var resp struct {
		Result []telegram.Update
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err)
	}
	return resp.Result, nil
}

func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) forwardCopyMany(to communications.Recipient, msgs []communications.Editable, key string, opts ...*communications.SendOptions) ([]telegram.Message, error) {
	params := map[string]string{
		"chat_id": to.Recipient(),
	}

	embedMessages(params, msgs)

	if len(opts) > 0 {
		b.RawEmbedSendOptions(params, opts[0])
	}

	data, err := b.Raw(key, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []Message
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		var resp struct {
			Result bool
		}
		if err := json.Unmarshal(data, &resp); err != nil {
			return nil, errors.Wrap(err)
		}
		return nil, errors.Wrap(err)
	}
	return resp.Result, nil
}

// extractOk checks given result for error. If result is ok returns nil.
// In other cases it extracts API error. If error is not presented
// in errors.go, it will be prefixed with `unknown` keyword.
func extractOk(data []byte) error {
	var e struct {
		Ok          bool                   `json:"ok"`
		Code        int                    `json:"error_code"`
		Description string                 `json:"description"`
		Parameters  map[string]interface{} `json:"parameters"`
	}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&e); err != nil {
		return fmt.Errorf("telebot: failed to decode response: %w", err)
	}
	if e.Ok {
		return nil
	}

	return errors.Wrap(fmt.Errorf("telegram: %s (%d)", e.Description, e.Code))
}
