package telegram

import (
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Venue represents a venue.
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    string   `json:"foursquare_id,omitempty"`
	FoursquareType  string   `json:"foursquare_type,omitempty"`
	GooglePlaceID   string   `json:"google_place_id,omitempty"`
	GooglePlaceType string   `json:"google_place_type,omitempty"`
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (v *Venue) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.AddFloat("latitude", v.Location.Lat)
	b.AddFloat("longitude", v.Location.Lng)
	b.Add("title", v.Title)
	b.Add("address", v.Address)
	b.Add("foursquare_id", v.FoursquareID)
	b.Add("foursquare_type", v.FoursquareType)
	b.Add("google_place_id", v.GooglePlaceID)
	b.Add("google_place_type", v.GooglePlaceType)

	return &outgoing.Method{
		Name:   "sendVenue",
		Params: b.Build(),
		Files:  nil,
	}
}
