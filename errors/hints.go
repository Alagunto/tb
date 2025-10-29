package errors

import (
	"fmt"
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

type HasRequestError struct {
	ErrorCode   int
	Description string
}

func (e *HasRequestError) Error() string {
	return fmt.Sprintf("telegram request error: %s (code: %d)", e.Description, e.ErrorCode)
}

func WithRequestError(err error, errorCode int, description string) error {
	return fmt.Errorf("%w: %w", err, &HasRequestError{ErrorCode: errorCode, Description: description})
}
