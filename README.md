# SwagBridge

## Description

SwagBridge is a command-line tool that converts Swagger/OpenAPI specifications (swagger.yaml) into Postman collections (postman_collection.json). It simplifies the process of generating API documentation and testing collections from your existing API specifications.

## Features

- Convert Swagger/OpenAPI specs to Postman collections
- Support for multiple API platforms:
  - [Insomnia](https://insomnia.rest/download)
  - [YAAK](https://yaak.app/)
  - [Paw](https://paw.cloud/)

## Installation

```bash
# Using npm
npm install -g swagbridge

# Using yarn
yarn global add swagbridge
```

## Usage

```bash
# Basic conversion
swagbridge convert -i swagger.yaml -o postman_collection.json

# Specify format (default is postman)
swagbridge convert -i swagger.yaml -o collection.json --format postman

# Available formats: postman, insomnia, yaak, paw
```

### Examples

1. Convert a local Swagger file:

```bash
swagbridge convert -i ./api/swagger.yaml -o ./collections/my_api.json
```

2. Convert from URL:

```bash
swagbridge convert -i https://api.example.com/swagger.yaml -o ./my_api_collection.json
```

3. Convert with specific format:

```bash
swagbridge convert -i swagger.yaml -o insomnia_collection.json --format insomnia
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).
