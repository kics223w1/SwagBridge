FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o swagbridge main.go

ENTRYPOINT ["./swagbridge"] 