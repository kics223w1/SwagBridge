package swagger

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

// SwaggerSpec represents the root structure of a Swagger/OpenAPI 2.0 specification
type SwaggerSpec struct {
	Swagger     string                 `json:"swagger" yaml:"swagger"`
	Info        Info                   `json:"info" yaml:"info"`
	Tags        []Tag                  `json:"tags" yaml:"tags"`
	Consumes    []string              `json:"consumes" yaml:"consumes"`
	Produces    []string              `json:"produces" yaml:"produces"`
	Paths       map[string]PathItem    `json:"paths" yaml:"paths"`
	Definitions map[string]Definition  `json:"definitions" yaml:"definitions"`
}

type Info struct {
	Title   string  `json:"title" yaml:"title"`
	Version string  `json:"version" yaml:"version"`
	Contact Contact `json:"contact" yaml:"contact"`
}

type Contact struct {
	Name  string `json:"name" yaml:"name"`
	URL   string `json:"url" yaml:"url"`
	Email string `json:"email" yaml:"email"`
}

type Tag struct {
	Name string `json:"name" yaml:"name"`
}

type PathItem struct {
	Get    *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Post   *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Patch  *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
}

type Operation struct {
	Summary     string       `json:"summary" yaml:"summary"`
	Description string       `json:"description" yaml:"description"`
	OperationID string       `json:"operationId" yaml:"operationId"`
	Parameters  []Parameter  `json:"parameters" yaml:"parameters"`
	Responses   map[string]Response `json:"responses" yaml:"responses"`
	Tags        []string     `json:"tags" yaml:"tags"`
}

type Parameter struct {
	Name     string      `json:"name" yaml:"name"`
	In       string      `json:"in" yaml:"in"`
	Required bool        `json:"required" yaml:"required"`
	Schema   *Schema     `json:"schema,omitempty" yaml:"schema,omitempty"`
	Type     string      `json:"type,omitempty" yaml:"type,omitempty"`
	Format   string      `json:"format,omitempty" yaml:"format,omitempty"`
}

type Response struct {
	Description string  `json:"description" yaml:"description"`
	Schema      *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

type Schema struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Definition struct {
	Type       string                 `json:"type" yaml:"type"`
	Properties map[string]Property    `json:"properties" yaml:"properties"`
}

type Property struct {
	Type   string  `json:"type" yaml:"type"`
	Format string  `json:"format,omitempty" yaml:"format,omitempty"`
	Ref    string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// ParseSwagger parses swagger content from YAML
func ParseSwagger(content []byte) (*SwaggerSpec, error) {
	var spec SwaggerSpec
	if err := yaml.Unmarshal(content, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}

// ParseSwaggerJSON parses swagger content from JSON
func ParseSwaggerJSON(content []byte) (*SwaggerSpec, error) {
	var spec SwaggerSpec
	if err := json.Unmarshal(content, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
} 