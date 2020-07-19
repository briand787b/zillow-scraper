package googlemaps_test

import (
	"context"
	"log"
	"os"
	"testing"
	"zcrapr/core/googlemaps"

	"googlemaps.github.io/maps"
)

var (
	mapClient *maps.Client
)

func TestMain(m *testing.M) {
	mapsKey := os.Getenv("GOOGLE_MAPS_BACKEND_API_KEY")
	if mapsKey == "" {
		log.Println("WARNING: GOOGLE_MAPS_BACKEND_API_KEY is empty string")
	}

	var err error
	mapClient, err = maps.NewClient(maps.WithAPIKey(mapsKey))
	if err != nil {
		log.Fatal("could not instantiate maps client: ", err)
	}
	os.Exit(m.Run())
}

func getCurrentGeocode(searchTerm string, t *testing.T) *maps.LatLng {
	geoCodeResults, err := mapClient.Geocode(context.Background(), &maps.GeocodingRequest{
		Address: searchTerm,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("lat long: ", geoCodeResults[0].Geometry.Location.String())
	return &geoCodeResults[0].Geometry.Location
}

func TestElevation(t *testing.T) {
	loc := getCurrentGeocode("577 Turkey Trot Rd, Clarkesville, GA 30532", t)
	results, err := mapClient.Elevation(context.Background(), &maps.ElevationRequest{
		Path: []maps.LatLng{
			*loc,
			maps.LatLng{
				Lat: loc.Lat + 0.1,
				Lng: loc.Lng,
			},
		},
		Samples: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	// t.Logf("elevations: %#v", results)
	t.Logf("number of samples: %d", len(results))
	var calculatedDistance float64
	var prevLatLng *maps.LatLng
	for i := 0; i < len(results); i++ {
		if prevLatLng != nil {
			distance := googlemaps.Distance(prevLatLng.Lat, prevLatLng.Lng, results[i].Location.Lat, results[i].Location.Lng, "K")
			// t.Logf("distance between points: %f", distance)
			calculatedDistance += distance
		}

		prevLatLng = results[i].Location
	}

	t.Logf("calculated distance: %f", calculatedDistance)

	t.Logf("starting LatLng: %#v", results[0].Location)
	t.Logf("ending LatLng: %#v", results[99].Location)
	totalDistance := googlemaps.Distance(results[0].Location.Lat, results[0].Location.Lng, results[99].Location.Lat, results[99].Location.Lng, "K")
	t.Logf("total distance: %f", totalDistance)
}

func TestPlacesAutocomplete(t *testing.T) {
	t.Skip()
	loc := getCurrentGeocode("445 Pippin Cir, Clarkesville, GA 30523", t)
	placeAutoResults, err := mapClient.PlaceAutocomplete(context.Background(), &maps.PlaceAutocompleteRequest{
		Input:    "Mt",
		Location: loc,
		Origin:   loc,
		Radius:   2000,
	})
	if err != nil {
		t.Fatal(err)
	}

	var predictionText string
	for _, prediction := range placeAutoResults.Predictions {
		predictionText = ""
		for _, term := range prediction.Terms {
			predictionText = predictionText + " " + term.Value
		}

		t.Log("term value: ", predictionText)
	}

	// t.Logf("result: %#v", result)
}

func TestFindPlaceFromText(t *testing.T) {
	loc := getCurrentGeocode("577 Turkey Trot Rd, Clarkesville, GA 30523", t)
	resp, err := mapClient.FindPlaceFromText(context.Background(), &maps.FindPlaceFromTextRequest{
		Input:             "restaurant",
		InputType:         maps.FindPlaceFromTextInputType("textquery"),
		LocationBiasPoint: loc,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}
