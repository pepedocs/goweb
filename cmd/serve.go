package main

import (
	"flag"
	"goweb/internal/server"
)

var (
	portFlag int
)

func main() {
	flag.IntVar(&portFlag, "port", 8000, "The tcp port to listen to")
	flag.Parse()
	server.Serve(portFlag)
}
