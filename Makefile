run:
	go run main.go -i swagger.yaml -h localhost:3000 -o postman_collection.json

build:
	go build -o swagger-to-postman main.go

.PHONY: run build