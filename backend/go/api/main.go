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
	cs, err := redis.NewCaptureRedisStore(l, "cap", "redis", "", 6379)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: grab these from env vars
	ps, err := redis.NewPropertyRedisStore(l, "id-counter", "id", "redis", "", 4, 1024, 6379)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(controller.Serve(*portFlag, l, cs, ps))
}
