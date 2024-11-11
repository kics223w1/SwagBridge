package main

import (
	"fmt"
	"llm-generate-test/swagger"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// Define the question for file path
	var filePath string
	prompt := &survey.Input{
		Message: "Enter the file path:",
		Help:    "Please provide the path to the file you want to read",
	}

	// Ask the question
	err := survey.AskOne(prompt, &filePath)
	if err != nil {
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

	// Print some basic information
	fmt.Printf("API Title: %s\n", spec.Info.Title)
	fmt.Printf("Version: %s\n", spec.Info.Version)
	fmt.Printf("Contact: %s (%s)\n", spec.Info.Contact.Name, spec.Info.Contact.Email)
	
	// Print available endpoints
	fmt.Println("\nAvailable Endpoints:")
	for path, item := range spec.Paths {
		if item.Get != nil {
			fmt.Printf("GET %s - %s\n", path, item.Get.Summary)
		}
		if item.Post != nil {
			fmt.Printf("POST %s - %s\n", path, item.Post.Summary)
		}
		if item.Patch != nil {
			fmt.Printf("PATCH %s - %s\n", path, item.Patch.Summary)
		}
	}
}
