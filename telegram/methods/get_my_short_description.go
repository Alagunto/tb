package methods

// GetMyShortDescriptionRequest represents the request for getMyShortDescription method.
type GetMyShortDescriptionRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyShortDescriptionResponse represents the response for getMyShortDescription method.
type GetMyShortDescriptionResponse struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
}
