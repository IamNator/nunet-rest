package main

import (
	"log"

	"nunet/app"
	"nunet/pkg"
)

const (
	defaultPort = 8080 // REST API port
)

func main() {
	// Read environment variables for configuration
	port := pkg.GetEnvOrDefaultInt("PORT", defaultPort)

	if err := app.Run(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
