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
go install github.com/kics223w1/swagbridge@latest
```

### Troubleshooting Installation

If you see `zsh: command not found: swagbridge` after installation, you need to add Go's bin directory to your PATH:

1. Add this line to your shell configuration file (`~/.zshrc` for zsh, `~/.bashrc` for bash):

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Reload your shell configuration:
   ```bash
   source ~/.zshrc  # for zsh
   # or
   source ~/.bashrc # for bash
   ```

Alternatively, you can run the binary directly using:

```bash
~/go/bin/swagbridge
```

## Usage

```bash
# Basic conversion
swagbridge -i api-specs/swagger.json -h localhost:3000 -s http -o collections/api_collection.json
```

### Examples

1. Convert a local Swagger file with host and scheme:

```bash
swagbridge -i ./api-specs/api.json -h api.example.com -s https -o ./collections/my_api.json
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).
