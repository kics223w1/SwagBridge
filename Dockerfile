# Build stage
FROM golang:1.23.3 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o swagbridge main.go

# Final stage
FROM alpine:latest
COPY --from=builder /app/swagbridge /app/swagbridge
ENTRYPOINT ["./swagbridge"] 