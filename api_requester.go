package tb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/files"
)

// ApiRequester is a struct that handles all application/json HTTP requests to the Telegram API
// For multipart/form-data, implement other requesters.
// TODO: implement multipart/form-data requester for attach:// schema file uploads
type ApiRequester[RequestSchema any, ResponseSchema any] struct {
	token          string
	apiURL         string
	client         *http.Client
	needsMultipart bool
	files          map[string]files.FileSource
}

func NewApiRequester[RequestSchema any, ResponseSchema any](token, apiURL string, client *http.Client) *ApiRequester[RequestSchema, ResponseSchema] {
	return &ApiRequester[RequestSchema, ResponseSchema]{
		token:          token,
		apiURL:         apiURL,
		client:         client,
		needsMultipart: false,
		files:          make(map[string]files.FileSource),
	}
}

func (r *ApiRequester[RequestSchema, ResponseSchema]) wrapRequestPreparationFailed(err error) error {
	return fmt.Errorf("%w: %w", err, errors.ErrRequestPreparationFailed)
}

func (r *ApiRequester[RequestSchema, ResponseSchema]) wrapRequestError(err error, req *http.Request, errorCode int, description string) error {
	err = errors.WithHttpRequest(err, req)
	if errorCode != 0 {
		err = errors.WithTelegramRequestErrorData(err, errorCode, description)
	}
	return err
}

func (r *ApiRequester[RequestSchema, ResponseSchema]) WithFileToUpload(fieldName string, file files.FileSource) *ApiRequester[RequestSchema, ResponseSchema] {
	r.needsMultipart = true
	r.files[fieldName] = file
	return r

}

func (r *ApiRequester[RequestSchema, ResponseSchema]) Request(ctx context.Context, method string, payload RequestSchema) (*ResponseSchema, error) {
	var contentType string
	var body io.Reader

	slog.Info("Request", "method", method, "payload", payload)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, r.wrapRequestPreparationFailed(err)
	}

	if r.needsMultipart {
		slog.Info("Request", "multipart", true)
		buf := bytes.NewBuffer(nil)
		multipartWriter := multipart.NewWriter(buf)

		// Note: file.Reader will be consumed here. If this ApiRequester is reused,
		// the file readers must be reset or recreated before the next request.
		for fieldName, file := range r.files {
			slog.Info("Request", "fieldName", fieldName, "file", file)
			fileWriter, err := multipartWriter.CreateFormFile(fieldName, file.GetFilenameForUpload())
			if err != nil {
				return nil, r.wrapRequestPreparationFailed(err)
			}
			if _, err = io.Copy(fileWriter, file.Reader); err != nil {
				return nil, r.wrapRequestPreparationFailed(err)
			}
		}

		jsonMap := map[string]any{}
		if err := json.Unmarshal(jsonPayload, &jsonMap); err != nil {
			return nil, r.wrapRequestPreparationFailed(err)
		}
		slog.Info("Request", "jsonMap", jsonMap)

		for fieldName, value := range jsonMap {
			slog.Info("Request", "fieldName", fieldName, "value", value)
			jsonnedValue, err := json.Marshal(value)
			if err != nil {
				return nil, r.wrapRequestPreparationFailed(err)
			}
			if err := multipartWriter.WriteField(fieldName, string(jsonnedValue)); err != nil {
				return nil, r.wrapRequestPreparationFailed(err)
			}
		}

		if err := multipartWriter.Close(); err != nil {
			return nil, r.wrapRequestPreparationFailed(err)
		}

		contentType = multipartWriter.FormDataContentType()
		body = buf
	} else {
		slog.Info("Request", "jsonPayload", string(jsonPayload))
		body = bytes.NewReader(jsonPayload)
		contentType = "application/json"
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.apiURL+"bot"+r.token+"/"+method, body)
	if err != nil {
		return nil, r.wrapRequestPreparationFailed(err)
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, r.wrapRequestError(err, req, 0, "")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, r.wrapRequestError(err, req, 0, "")
	}

	/*
		The response contains a JSON object, which always has a Boolean field 'ok' and may have an optional String field 'description' with a human-readable description of the result.
		If 'ok' equals True, the request was successful and the result of the query can be found in the 'result' field.
		In case of an unsuccessful request, 'ok' equals false and the error is explained in the 'description'.
		An Integer 'error_code' field is also returned, but its contents are subject to change in the future.
		Some errors may also have an optional field 'parameters' of the type ResponseParameters, which can help to automatically handle the error.
	*/
	var response struct {
		Result      ResponseSchema `json:"result"`
		Ok          bool           `json:"ok"`
		Description string         `json:"description"`
		ErrorCode   int            `json:"error_code"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, r.wrapRequestError(err, req, response.ErrorCode, response.Description)
	}

	if !response.Ok {
		return nil, r.wrapRequestError(errors.ErrRequestFailed,
			req,
			response.ErrorCode,
			response.Description,
		)
	}
	return &response.Result, nil
}
