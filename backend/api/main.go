package main

import (
	"flag"
	"log"
	"os"

	"zcrape/api/controller"
	"zcrape/core/plog"

	"github.com/google/uuid"
)

var (
	portFlag = flag.Int("port", 0, "the port to listen on")
)

func main() {
	flag.Parse()

	log.Printf("serving from port %v\n", *portFlag)

	l := plog.NewPLogger(log.New(os.Stdout, "", 0), uuid.New())

	controller.Serve(*portFlag, l, nil)
}
