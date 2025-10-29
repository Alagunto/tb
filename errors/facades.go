package errors

func BuildRequestFailedError(errorCode int, description string) error {
	return WithRequestError(ErrRequestFailed, errorCode, description)
}
