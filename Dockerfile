FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go module files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN go build -o /app/metrics-typer ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/metrics-typer /app/metrics-typer

# Set default environment variables
ENV LISTEN_ADDR=":8080"
ENV DESTINATION_URL=""

CMD ["/app/metrics-typer"]
