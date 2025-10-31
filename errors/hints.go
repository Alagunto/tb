package errors

import (
	"fmt"
	"net/http"
)

type InvalidParamError struct {
	ParamName  string
	ParamValue any
}

func (e *InvalidParamError) Error() string {
	return fmt.Sprintf("invalid parameter: %s=%v", e.ParamName, e.ParamValue)
}

func WithInvalidParam(err error, paramName string, paramValue any) error {
	return fmt.Errorf("%w: %w", err, &InvalidParamError{ParamName: paramName, ParamValue: paramValue})
}

type MissingEntity string

const (
	MissingEntityMessage       MissingEntity = "message"
	MissingEntityCallback      MissingEntity = "callback"
	MissingEntityInlineQuery   MissingEntity = "inline_query"
	MissingEntityShippingQuery MissingEntity = "shipping_query"
)

type MissingEntityError struct {
	MissingEntity MissingEntity
}

func (e *MissingEntityError) Error() string {
	return fmt.Sprintf("context is missing: %s", e.MissingEntity)
}

func WithMissingEntity(err error, missingEntity MissingEntity) error {
	return fmt.Errorf("%w: %w", err, &MissingEntityError{MissingEntity: missingEntity})
}

type HasTelegramRequestErrorData struct {
	ErrorCode   int
	Description string
}

func (e *HasTelegramRequestErrorData) Error() string {
	return fmt.Sprintf("telegram request error: %s (code: %d)", e.Description, e.ErrorCode)
}

func WithTelegramRequestErrorData(err error, errorCode int, description string) error {
	return fmt.Errorf("%w: %w", err, &HasTelegramRequestErrorData{ErrorCode: errorCode, Description: description})
}

type HasHttpRequest struct {
	Request *http.Request
}

func (e *HasHttpRequest) Error() string {

	return fmt.Sprintf("(url: %s, method: %s)", e.Request.URL.String(), e.Request.Method)
}

func WithHttpRequest(err error, request *http.Request) error {
	return fmt.Errorf("%w: %w", err, &HasHttpRequest{Request: request})
}
