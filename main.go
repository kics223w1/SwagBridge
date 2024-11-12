package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"llm-generate-test/postman"
	"llm-generate-test/swagger"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// Define the question for file path
	var filePath string
	var host string

	prompt := &survey.Input{
		Message: "Enter the file path:",
		Help:    "Please provide the path to the file you want to read",
	}

	promptHost := &survey.Input{
		Message: "Enter the host:",
		Help:    "Please provide the host for the collection",
	}

	// Ask the question
	err := survey.AskOne(prompt, &filePath)
	err2 := survey.AskOne(promptHost, &host)
	if err != nil || err2 != nil {
		fmt.Printf("Error getting input: %v\n", err)
		return
	}

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse the swagger specification
	spec, err := swagger.ParseSwagger(content)
	if err != nil {
		fmt.Printf("Error parsing swagger: %v\n", err)
		return
	}

	// Generate Postman collection
	collection, err := postman.GeneratePostmanCollection(spec , host)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(collection, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Write to file
	err = ioutil.WriteFile("postman_collection.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
