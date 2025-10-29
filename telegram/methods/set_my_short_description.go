package methods

// SetMyShortDescriptionRequest represents the request for setMyShortDescription method.
type SetMyShortDescriptionRequest struct {
	// New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	ShortDescription string `json:"short_description,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyShortDescriptionResponse represents the response for setMyShortDescription method.
type SetMyShortDescriptionResponse bool
