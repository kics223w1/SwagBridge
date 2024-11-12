run:
	go run main.go -i ./api-specs/server1.json -h localhost:3000 -s http -o ./collections/server1.json

build:
	GOOS=darwin GOARCH=arm64 go build -o swagbridge main.go

docker-build:
	docker build -t swagbridge .

.PHONY: run build docker-build