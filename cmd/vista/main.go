package main

import (
	"flag"
	"log"
	"os"

	"github.com/jamie/vista/api"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port for the API server")
	flag.Parse()

	server := api.NewServer(port)
	if err := server.Start(); err != nil {
		log.Printf("Server error: %v", err)
		os.Exit(1)
	}
}
