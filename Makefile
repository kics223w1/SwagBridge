run:
	go run main.go -i ./api-specs/yaml/simple-bank.yaml -h simplebank.com -o ./collections/simplebank_collection.json

build:
	GOOS=darwin GOARCH=arm64 go build -o swagbridge main.go

docker-build:
	docker build -t swagbridge .

.PHONY: run build docker-build