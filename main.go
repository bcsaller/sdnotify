package main

import (
	"flag"
	"log"
	"strings"

	"github.com/okzk/sdnotify"
)

func main() {
	flag.Parse()
	args := flag.Args()
	switch cmd := args[0]; cmd {
	case "ready":
		err := sdnotify.Ready()
		if err != nil {
			log.Fatal(err)
		}

	case "watchdog":
		err := sdnotify.Watchdog()
		if err != nil {
			log.Fatal(err)
		}

	default:
		err := sdnotify.Status(strings.Join(args, " "))
		if err != nil {
			log.Fatal(err)
		}

	}

}
