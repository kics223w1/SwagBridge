package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"llm-generate-test/postman"
	"llm-generate-test/swagger"
	"log"
	"os"
)

func main() {
	// Define command line flags
	inputFile := flag.String("i", "", "Path to the input Swagger/OpenAPI YAML file")
	schema := flag.String("s", "https", "Schema to use (http, https, ws, wss)")
	host := flag.String("h", "", "Host for the API endpoints")
	outputFile := flag.String("o", "postman_collection.json", "Output file path for the Postman collection")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if *inputFile == "" || *host == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Read file content
	content, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	// Parse the swagger specification
	spec, err := swagger.ParseSwagger(content)
	if err != nil {
		log.Fatalf("Error parsing swagger: %v\n", err)
	}

	// Generate Postman collection
	collection, err := postman.GeneratePostmanCollection(spec, *host, *schema)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(collection, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Write to file
	err = ioutil.WriteFile(*outputFile, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully generated Postman collection: %s\n", *outputFile)
}
