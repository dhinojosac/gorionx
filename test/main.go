package main

// Import packets orionxctools.
import (
	"fmt"
	"log"
	"os"

	orionxctools "github.com/dhinojosac/gorionx"
)

// The KEY and API KEY were stored in environment variables.

func main() {
	//Read environmet variables
	fmt.Println("Reading Environment Variables")
	var key string
	var apiKey string

	key, ok := os.LookupEnv("ORIONX_KEY")
	if !ok {
		log.Fatal("There are not present ORIONX_KEY env var\n")
	}
	apiKey, ok = os.LookupEnv("ORIONX_API_KEY")
	if !ok {
		log.Fatal("There are not present ORIONX_API_KEY env var\n")
	}
	fmt.Printf("api: %s \napi_key: %s", key, apiKey)

	// Generate the queries you need
	queries := []string{
		`{market(code:"CHACLP"){
			lastTrade{
				price
			}
		  }}`,
		`{cha:marketOrderBook (marketCode:"CHACLP") {
			spread
			mid
			}}`,
	}

	for _, query := range queries {
		orionxctools.MakeRequest(key, apiKey, query) // Performs API requests
	}
}
