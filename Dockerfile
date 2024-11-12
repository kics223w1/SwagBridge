# Build stage
FROM golang:1.23.3 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 go build -o swagbridge main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/swagbridge /app/swagbridge
RUN chmod +x /app/swagbridge
ENTRYPOINT ["/app/swagbridge"]