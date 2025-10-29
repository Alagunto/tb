package methods

// SetMyDescriptionRequest represents the request for setMyDescription method.
type SetMyDescriptionRequest struct {
	// New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	Description string `json:"description,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyDescriptionResponse represents the response for setMyDescription method.
type SetMyDescriptionResponse bool
