package postman

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

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
	Event    []Event  `json:"event,omitempty"`
}

type Event struct {
	Listen string    `json:"listen"`
	Script EventScript `json:"script"`
}

type EventScript struct {
	Type     string   `json:"type"`
	Exec     []string `json:"exec"`
	Packages struct{} `json:"packages,omitempty"`
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
func GeneratePostmanCollection(spec *swagger.SwaggerSpec, host string, schema string) (*Collection, error) {
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
		items := convertPathToItems(path, pathItem, host, schema)
		collection.Item = append(collection.Item, items...)
	}

	return collection, nil
}

func convertPathToItems(path string, pathItem swagger.PathItem, host string, schema string) []Item {
	items := make([]Item, 0)

	// Handle GET operations
	if pathItem.Get != nil {
		items = append(items, createItem("GET", path, pathItem.Get, host, schema))
	}

	// Handle POST operations
	if pathItem.Post != nil {
		items = append(items, createItem("POST", path, pathItem.Post, host, schema))
	}

	// Handle PATCH operations
	if pathItem.Patch != nil {
		items = append(items, createItem("PATCH", path, pathItem.Patch, host, schema))
	}

	return items
}

func createItem(method, path string, operation *swagger.Operation, host string, schema string) Item {
	item := Item{
		Name: operation.Summary,
		Request: Request{
			Method: method,
			URL:    createRequestURL(path, operation.Parameters, host, schema),
		},
		Response: make([]string, 0),
	}

	// Add test script for status code validation
	if operation.Responses != nil {
		testScript := createStatusCodeTestScript(operation.Responses)
		item.Event = []Event{
			{
				Listen: "test",
				Script: EventScript{
					Type: "text/javascript",
					Exec: []string{testScript},
					Packages: struct{}{},
				},
			},
		}
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

func createRequestURL(path string, parameters []swagger.Parameter, host string, schema string) RequestURL {
	url := RequestURL{
		Protocol: schema,
		Host:     []string{host},
		Path:     strings.Split(strings.Trim(path, "/"), "/"),
		Query:    make([]URLQuery, 0),
	}

	// Add query parameters with random values based on type
	for _, param := range parameters {
		if param.In == "query" {
			randomValue := generateRandomValueByType(param.Type, param.Format)
			url.Query = append(url.Query, URLQuery{
				Key:   param.Name,
				Value: randomValue,
			})
		}
	}

	// Construct raw URL using the provided schema
	raw := fmt.Sprintf("%s://%s%s", schema, host, path)
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

// New helper function to generate random values
func generateRandomValueByType(paramType, format string) string {
	switch paramType {
	case "string":
		switch format {
		case "uuid":
			return uuid.New().String()
		case "date-time":
			return time.Now().Format(time.RFC3339)
		case "email":
			return fmt.Sprintf("user%d@example.com", rand.Intn(1000))
		default:
			return fmt.Sprintf("sample_string_%d", rand.Intn(1000))
		}
	case "integer", "number":
		return fmt.Sprintf("%d", rand.Intn(100))
	case "boolean":
		return fmt.Sprintf("%v", rand.Intn(2) == 1)
	default:
		return "sample_value"
	}
}

func createStatusCodeTestScript(responses map[string]swagger.Response) string {
	var validCodes []string
	for code := range responses {
		validCodes = append(validCodes, code)
	}
	
	script := `
pm.test("Status code is expected", function() {
    var expectedCodes = [%s];
    pm.expect(expectedCodes).to.include(pm.response.code.toString());
});`

	// Convert status codes to comma-separated string of quoted values
	quotedCodes := make([]string, len(validCodes))
	for i, code := range validCodes {
		quotedCodes[i] = fmt.Sprintf(`"%s"`, code)
	}
	
	return fmt.Sprintf(script, strings.Join(quotedCodes, ", "))
} 