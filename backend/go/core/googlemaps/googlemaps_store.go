package googlemaps

import (
	"googlemaps.github.io/maps"
)

// GoogleMapStore is a MapStore backed by GoogleMaps
type GoogleMapStore struct {
	client *maps.Client
}

// NewGoogleMapStore creates a new MapStore backed by GoogleMaps
func NewGoogleMapStore(client *maps.Client) *GoogleMapStore {
	return &GoogleMapStore{client: client}
}

// func (s *GoogleMapStore)
