package methods

// GetMyNameRequest represents the request for getMyName method.
type GetMyNameRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyNameResponse represents the response for getMyName method.
type GetMyNameResponse struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
}
