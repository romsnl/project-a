package main

import (
	"flag"

	"github.com/romsnl/project-a/api"
)

var (
	laddr = flag.String("listenaddr", ":3001", "Server listen address")
)

func main() {
	flag.Parse()

	api := api.NewServer()
	api.Serve(*laddr)
}
