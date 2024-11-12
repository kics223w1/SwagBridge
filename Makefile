run:
	go run main.go -i ./api-specs/localhost-server.json -h localhost:3000 -s http -o ./collections/simplebank_collection.json

build:
	GOOS=darwin GOARCH=arm64 go build -o swagbridge main.go

docker-build:
	docker build -t swagbridge .

.PHONY: run build docker-build