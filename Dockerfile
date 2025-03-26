FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git protobuf-dev

# Copy go module files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Generate protobuf code
COPY internal/proto/ ./internal/proto/
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN protoc --go_out=. --go_opt=paths=source_relative internal/proto/remote.proto

# Copy source code
COPY . .

# Build application
RUN go build -o /app/metrics-transformer ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/metrics-transformer /app/metrics-transformer

# Set default environment variables
ENV LISTEN_ADDR=":8080"
ENV DESTINATION_URL=""

CMD ["/app/metrics-transformer"]
