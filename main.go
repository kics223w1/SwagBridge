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
	var filePath, host, schema string

	// Define questions
	prompt := &survey.Input{
		Message: "Enter the YAML file path:",
		Help:    "Please provide the path to the YAML file you want to read",
	}

	promptHost := &survey.Input{
		Message: "Enter the host:",
		Help:    "Please provide the host for the collection",
	}

	// Define schema options
	schemaOptions := []string{
		"http",
		"https",
		"ws",
		"wss",
		"Enter manually",
	}

	promptSchema := &survey.Select{
		Message: "Choose the schema:",
		Options: schemaOptions,
		Help:    "Select a schema or choose 'Enter manually' to input your own",
	}

	// Ask the questions
	err := survey.AskOne(prompt, &filePath)
	if err != nil {
		fmt.Printf("Error getting file path: %v\n", err)
		return
	}

	err = survey.AskOne(promptHost, &host)
	if err != nil {
		fmt.Printf("Error getting host: %v\n", err)
		return
	}

	err = survey.AskOne(promptSchema, &schema)
	if err != nil {
		fmt.Printf("Error getting schema: %v\n", err)
		return
	}

	// If user chose to enter manually, prompt for custom schema
	if schema == "Enter manually" {
		promptCustomSchema := &survey.Input{
			Message: "Enter your custom schema:",
			Help:    "Please provide the schema (e.g., ftp, sftp)",
		}
		err = survey.AskOne(promptCustomSchema, &schema)
		if err != nil {
			fmt.Printf("Error getting custom schema: %v\n", err)
			return
		}
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

	// Generate Postman collection (you'll need to update this function to accept schema)
	collection, err := postman.GeneratePostmanCollection(spec, host, schema)
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
