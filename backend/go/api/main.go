package main

import (
	"flag"
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
	portEnvVar             = "PORT"
	redisHostEnvVar        = "REDIS_HOST"
	googleMapsAPIKeyEnvVar = "GOOGLE_MAPS_API_KEY"
)

var (
	portFlag = flag.Int("port", 0, "the port to listen on")
)

func main() {
	flag.Parse()

	log.Printf("serving from port %v\n", *portFlag)

	l := plog.NewPLogger(log.New(os.Stdout, "", 0), uuid.New())
	cs := getCaptureRedisStore(l)
	ps := getPropertyRedisStore(l)

	log.Fatalln(controller.Serve(getPort(), l, cs, ps))
}

func getCaptureRedisStore(l plog.Logger) *redis.CaptureRedisStore {
	redisHost := getRedisHost()
	cs, err := redis.NewCaptureRedisStore(l, "cap", redisHost, "", 6379)
	if err != nil {
		log.Fatalln(err)
	}

	return cs
}

func getGoogleMapsStore() *googlemaps.GoogleMapStore {
	mapsAPIKey := os.Getenv(googleMapsAPIKeyEnvVar)
	if mapsAPIKey == "" {
		log.Printf("WARNING: %s is emtpy string", googleMapsAPIKeyEnvVar)
	}

	client, err := maps.NewClient(maps.WithAPIKey(mapsAPIKey))
	if err != nil {
		log.Fatalf("ERROR: could not get google maps client by key %s\n", googleMapsAPIKeyEnvVar)
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
