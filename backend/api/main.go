package main

import (
	"flag"
	"log"
	"os"

	"zcrapr/api/controller"
	"zcrapr/core/plog"
	"zcrapr/core/redis"

	"github.com/google/uuid"
)

var (
	portFlag = flag.Int("port", 0, "the port to listen on")
)

func main() {
	flag.Parse()

	log.Printf("serving from port %v\n", *portFlag)

	l := plog.NewPLogger(log.New(os.Stdout, "", 0), uuid.New())

	// TODO: grab these from env vars
	rc, err := redis.NewPropertyRedisStore(l, "id-counter", "id", "cap", "redis", "", 4, 6379)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(controller.Serve(*portFlag, l, rc))
}
