package postman

import (
	"fmt"
	"strings"

	"llm-generate-test/swagger"

	"github.com/google/uuid"
)

// Collection represents a Postman collection structure
type Collection struct {
	Info     Info       `json:"info"`
	Item     []Item     `json:"item"`
	Variable []Variable `json:"variable,omitempty"`
}

type Info struct {
	PostmanID string `json:"_postman_id"`
	Name      string `json:"name"`
	Schema    string `json:"schema"`
}

type Item struct {
	Name     string   `json:"name"`
	Request  Request  `json:"request"`
	Response []string `json:"response"`
}

type Request struct {
	Method string      `json:"method"`
	Header []Header    `json:"header,omitempty"`
	URL    RequestURL  `json:"url"`
	Body   *RequestBody `json:"body,omitempty"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RequestURL struct {
	Raw      string      `json:"raw"`
	Protocol string      `json:"protocol"`
	Host     []string    `json:"host"`
	Path     []string    `json:"path"`
	Query    []URLQuery  `json:"query,omitempty"`
}

type URLQuery struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RequestBody struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw,omitempty"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GeneratePostmanCollection converts a Swagger spec to a Postman collection
func GeneratePostmanCollection(spec *swagger.SwaggerSpec) (*Collection, error) {
	collection := &Collection{
		Info: Info{
			PostmanID: uuid.New().String(),
			Name:      spec.Info.Title,
			Schema:    "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Item: make([]Item, 0),
	}

	// Convert each path to Postman items
	for path, pathItem := range spec.Paths {
		items := convertPathToItems(path, pathItem)
		collection.Item = append(collection.Item, items...)
	}

	return collection, nil
}

func convertPathToItems(path string, pathItem swagger.PathItem) []Item {
	items := make([]Item, 0)

	// Handle GET operations
	if pathItem.Get != nil {
		items = append(items, createItem("GET", path, pathItem.Get))
	}

	// Handle POST operations
	if pathItem.Post != nil {
		items = append(items, createItem("POST", path, pathItem.Post))
	}

	// Handle PATCH operations
	if pathItem.Patch != nil {
		items = append(items, createItem("PATCH", path, pathItem.Patch))
	}

	return items
}

func createItem(method, path string, operation *swagger.Operation) Item {
	item := Item{
		Name: operation.Summary,
		Request: Request{
			Method: method,
			URL:    createRequestURL(path, operation.Parameters),
		},
		Response: make([]string, 0),
	}

	// Add headers if needed
	if method == "POST" || method == "PATCH" {
		item.Request.Header = []Header{
			{
				Key:   "Content-Type",
				Value: "application/json",
			},
		}
	}

	// Add request body for POST/PATCH methods
	if method == "POST" || method == "PATCH" {
		item.Request.Body = &RequestBody{
			Mode: "raw",
			Raw:  "{\n    \"key\": \"value\"\n}", // You might want to generate this based on the schema
		}
	}

	return item
}

func createRequestURL(path string, parameters []swagger.Parameter) RequestURL {
	url := RequestURL{
		Protocol: "https",
		Host:     []string{"{{baseUrl}}"},
		Path:     strings.Split(strings.Trim(path, "/"), "/"),
		Query:    make([]URLQuery, 0),
	}

	// Add query parameters
	for _, param := range parameters {
		if param.In == "query" {
			url.Query = append(url.Query, URLQuery{
				Key:   param.Name,
				Value: fmt.Sprintf("{{%s}}", param.Name),
			})
		}
	}

	// Construct raw URL
	raw := fmt.Sprintf("{{baseUrl}}%s", path)
	if len(url.Query) > 0 {
		raw += "?"
		queries := make([]string, 0)
		for _, q := range url.Query {
			queries = append(queries, fmt.Sprintf("%s=%s", q.Key, q.Value))
		}
		raw += strings.Join(queries, "&")
	}
	url.Raw = raw

	return url
} 