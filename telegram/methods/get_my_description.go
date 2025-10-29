package methods

// GetMyDescriptionRequest represents the request for getMyDescription method.
type GetMyDescriptionRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyDescriptionResponse represents the response for getMyDescription method.
type GetMyDescriptionResponse struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
}
