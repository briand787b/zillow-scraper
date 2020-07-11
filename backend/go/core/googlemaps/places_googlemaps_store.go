package googlemaps

import (
	"googlemaps.github.io/maps"
)

type PlacesGoogleMapsStore struct {
	client *maps.Client
}

func NewPlacesGoogleMapsStore(client *maps.Client) *PlacesGoogleMapsStore {
	return &PlacesGoogleMapsStore{client: client}
}
