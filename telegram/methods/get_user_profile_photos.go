package methods

import "github.com/alagunto/tb/telegram"

// GetUserProfilePhotosRequest represents the request for getUserProfilePhotos method.
type GetUserProfilePhotosRequest struct {
	// Unique identifier of the target user
	UserID string `json:"user_id"`

	// Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset int `json:"offset,omitempty"`

	// Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
}

// GetUserProfilePhotosResponse represents the response for getUserProfilePhotos method.
type GetUserProfilePhotosResponse struct {
	TotalCount int              `json:"total_count"`
	Photos     []telegram.Photo `json:"photos"`
}
