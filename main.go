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
	"path/filepath"
	"strings"
)

func main() {
	// Define command line flags
	inputFile := flag.String("i", "", "Path to the input JSON file")
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

	// Ensure output file has .json extension
	if !strings.HasSuffix(*outputFile, ".json") {
		*outputFile = *outputFile + ".json"
	}

	// If output path doesn't contain directory, use current directory
	outputPath := *outputFile
	if !strings.Contains(outputPath, string(os.PathSeparator)) {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current directory: %v\n", err)
		}
		outputPath = filepath.Join(currentDir, outputPath)
	}

	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v\n", err)
	}

	// Read file content
	content, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	// Validate JSON format
	if !json.Valid(content) {
		log.Fatalf("Error: Input file is not valid JSON\n")
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
	err = ioutil.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully generated Postman collection at: %s\n", outputPath)
}
