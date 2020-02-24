package main

import (
	"flag"
	"log"
	"strconv"
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

	case "errno":
		errno, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		err = sdnotify.Errno(errno)
		if err != nil {
			log.Fatal(err)
		}

	case "stopping":
		err := sdnotify.Stopping()
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
