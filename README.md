# Metrics Transformer

A lightweight service that receives Prometheus remote_write metrics, transforms them by modifying metric types based on name patterns, and forwards them to another Prometheus-compatible endpoint.

## Features

- Receives metrics via Prometheus remote_write protocol
- Transforms metric types based on hardcoded name patterns
- Forwards transformed metrics to a configurable destination
- Minimal dependencies (only protobuf and snappy compression libraries)

## Metric Type Mapping

The service uses a hardcoded mapping of metric name patterns to desired types:

| Pattern         | New Type   |
|-----------------|------------|
| latency         | HISTOGRAM  |
| count           | COUNTER    |
| memory          | GAUGE      |
| bytes           | GAUGE      |
| cpu             | GAUGE      |
| requests_total  | COUNTER    |
| errors_total    | COUNTER    |
| duration_seconds| HISTOGRAM  |
| response_time   | HISTOGRAM  |
| connections     | GAUGE      |
| temperature     | GAUGE      |
| uptime          | COUNTER    |
| queue_length    | GAUGE      |
| pending_tasks   | GAUGE      |
| resource_usage  | GAUGE      |
| throughput      | GAUGE      |
| operations_total| COUNTER    |
| failures_total  | COUNTER    |
| success_ratio   | GAUGE      |
| processing_time | HISTOGRAM  |

This mapping is defined in `internal/config/config.go` and can be modified directly in the code.

## Configuration

Configuration is provided via environment variables:

- `LISTEN_ADDR`: Address to listen on (default: ":8080")
- `DESTINATION_URL`: URL to forward metrics to (required)

## Usage

```bash
# Run with default listen address
DESTINATION_URL="http://prometheus:9090/api/v1/write" ./metrics-transformer

# Run with custom listen address
LISTEN_ADDR=":9090" DESTINATION_URL="http://prometheus:9090/api/v1/write" ./metrics-transformer
```

## Building

```bash
# Generate protobuf code
protoc --go_out=. --go_opt=paths=source_relative internal/proto/remote.proto

# Build binary
go build -o metrics-transformer ./cmd/server
```

## Docker

```bash
# Build image
docker build -t metrics-transformer .

# Run container
docker run -p 8080:8080 -v $(pwd)/config:/app/config metrics-transformer
```