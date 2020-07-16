package main

import (
	"log"
	"os"
	"strconv"

	"zcrapr/api/controller"
	"zcrapr/core/googlemaps"
	"zcrapr/core/plog"
	"zcrapr/core/redis"

	"github.com/google/uuid"
	"googlemaps.github.io/maps"
)

const (
	portEnvVar                    = "PORT"
	redisHostEnvVar               = "REDIS_HOST"
	googleMapsBackendAPIKeyEnvVar = "GOOGLE_MAPS_BACKEND_API_KEY"
	googleMapsEmbedAPIKeyEnvVar   = "GOOGLE_MAPS_EMBED_API_KEY"
)

func main() {
	l := plog.NewPLogger(log.New(os.Stdout, "", 0), uuid.New())
	cs := getCaptureRedisStore(l)
	ps := getPropertyRedisStore(l)
	getGoogleMapsStore(getGoogleMapsBackendAPIKey())
	embedKey := getGoogleMapsEmbedAPIKey()
	port := getPort()

	log.Fatalln(controller.Serve(port, l, embedKey, cs, ps))
}

func getCaptureRedisStore(l plog.Logger) *redis.CaptureRedisStore {
	redisHost := getRedisHost()
	cs, err := redis.NewCaptureRedisStore(l, "cap", redisHost, "", 6379)
	if err != nil {
		log.Fatalln(err)
	}

	return cs
}

func getGoogleMapsBackendAPIKey() string {
	mapsAPIKey := os.Getenv(googleMapsBackendAPIKeyEnvVar)
	if mapsAPIKey == "" {
		log.Printf("WARNING: %s is emtpy string", googleMapsBackendAPIKeyEnvVar)
	}

	return mapsAPIKey
}

func getGoogleMapsEmbedAPIKey() string {
	mapsAPIKey := os.Getenv(googleMapsEmbedAPIKeyEnvVar)
	if mapsAPIKey == "" {
		log.Printf("WARNING: %s is emtpy string", googleMapsEmbedAPIKeyEnvVar)
	}

	return mapsAPIKey
}

func getGoogleMapsStore(googleMapsBackendAPIKey string) *googlemaps.GoogleMapStore {
	client, err := maps.NewClient(maps.WithAPIKey(googleMapsBackendAPIKey))
	if err != nil {
		log.Fatalf("ERROR: could not get google maps client by key %s: %s\n", googleMapsBackendAPIKeyEnvVar, err)
	}

	return googlemaps.NewGoogleMapStore(client)
}

func getPort() int {
	portStr := os.Getenv(portEnvVar)
	if portStr == "" {
		log.Printf("WARNING: %s is empty string\n", portEnvVar)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	return port
}

func getPropertyRedisStore(l plog.Logger) *redis.PropertyRedisStore {
	redisHost := getRedisHost()
	ps, err := redis.NewPropertyRedisStore(l, "id-counter", "id", redisHost, "", 4, 1024, 6379)
	if err != nil {
		log.Fatalln(err)
	}

	return ps
}

func getRedisHost() string {
	redisHost := os.Getenv(redisHostEnvVar)
	if redisHost == "" {
		log.Printf("WARNING: %s is empty string\n", redisHostEnvVar)
	}

	return redisHost
}
