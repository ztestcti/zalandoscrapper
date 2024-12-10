#maintainer max.mustermann@zalando.de

package main

import (
	"log"
	"net/http"
	"aws-mock-app/config"
	"aws-mock-app/handlers"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/aws", handlers.AWSHandler)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
