package main

import (
	"fmt"
	"github.com/antonybudianto/resto-api/app"
	"log"
	"os"
	"strconv"
)

func getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func main() {
	apiPortStr := getenv("API_PORT", "8000")
	apiPort, err := strconv.Atoi(apiPortStr)

	if err != nil {
		log.Fatal(fmt.Sprintf("API_PORT should be number. Received: %s", apiPortStr))
		os.Exit(2)
	}

	log.Printf("Resto API running at %d", apiPort)

	a := app.App{}
	a.Initialize("root", "hello", "rest_api_example")
	a.Run(fmt.Sprintf(":%d", apiPort))
}
