package methods

// SetMyNameRequest represents the request for setMyName method.
type SetMyNameRequest struct {
	// New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	Name string `json:"name,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyNameResponse represents the response for setMyName method.
type SetMyNameResponse bool
