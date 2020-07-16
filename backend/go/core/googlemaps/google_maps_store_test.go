package googlemaps_test

import (
	"log"
	"os"
	"testing"
)

var (
	string mapsAPIKey
)

func TestMain(m *testing.M) {
	mapsAPIKey = os.Getenv("GOOGLE_MAPS_API_KEY")
	if mapsAPIKey == "" {
		log.Println("WARNING: GOOGLE_MAPS_API_KEY is empty string")
	}
	os.Exit(m.Run())
}
