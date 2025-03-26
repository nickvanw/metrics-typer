package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang/snappy"
	"github.com/planetscale/metrics-typer/internal/config"
	promproto "github.com/planetscale/metrics-typer/internal/proto"
	"github.com/planetscale/metrics-typer/internal/transformer"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	Config      *config.Config
	Transformer *transformer.Transformer
}

func New(cfg *config.Config) *Server {
	return &Server{
		Config:      cfg,
		Transformer: transformer.New(),
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/receive", s.handleRemoteWrite)
	log.Printf("Starting server on %s", s.Config.ListenAddr)
	return http.ListenAndServe(s.Config.ListenAddr, nil)
}

func (s *Server) handleRemoteWrite(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received remote_write request from %s", r.RemoteAddr)

	if r.Method != http.MethodPost {
		log.Printf("Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the compressed data
	compressed, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	// Decompress the data using snappy
	decompressed, err := snappy.Decode(nil, compressed)
	if err != nil {
		log.Printf("Error decompressing request body: %v", err)
		http.Error(w, fmt.Sprintf("Error decompressing request body: %v", err), http.StatusBadRequest)
		return
	}

	// Parse protobuf data
	var req promproto.WriteRequest
	if err := proto.Unmarshal(decompressed, &req); err != nil {
		log.Printf("Error unmarshaling protobuf: %v", err)
		http.Error(w, fmt.Sprintf("Error unmarshaling protobuf: %v", err), http.StatusBadRequest)
		return
	}

	// Log detailed information about the received metrics
	log.Printf("Parsed WriteRequest with %d series and %d metadata entries", len(req.Timeseries), len(req.Metadata))

	if len(req.Timeseries) > 0 {
		log.Printf("Timeseries details (showing up to %d):", 10)
		for i, ts := range req.Timeseries {
			if i >= 10 {
				log.Printf("... and %d more timeseries", len(req.Timeseries)-10)
				break
			}

			// Build label string
			var labelStr string
			for _, label := range ts.Labels {
				if labelStr != "" {
					labelStr += ", "
				}
				labelStr += fmt.Sprintf("%s=%s", label.Name, label.Value)
			}

			// Log samples information
			sampleInfo := fmt.Sprintf("samples=%d", len(ts.Samples))
			if len(ts.Samples) > 0 {
				// Show first sample timestamp and value
				firstSample := ts.Samples[0]
				timestamp := time.Unix(0, firstSample.Timestamp*1000000) // Convert ms to ns
				sampleInfo += fmt.Sprintf(", first_value=%.4f, timestamp=%s",
					firstSample.Value, timestamp.Format(time.RFC3339))
			}

			log.Printf("TIMESERIES[%d]: labels={%s}, %s",
				i, labelStr, sampleInfo)
		}
	}

	// Transform metrics (with logging in the transformer)
	s.Transformer.Transform(&req)

	// Reserialize the transformed data
	reserializedProto, err := proto.Marshal(&req)
	if err != nil {
		log.Printf("Error marshaling protobuf: %v", err)
		http.Error(w, fmt.Sprintf("Error marshaling protobuf: %v", err), http.StatusInternalServerError)
		return
	}

	// Forward to destination
	if err := s.forwardMetrics(snappy.Encode(nil, reserializedProto), r); err != nil {
		log.Printf("Error forwarding metrics: %v", err)
		http.Error(w, fmt.Sprintf("Error forwarding metrics: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("Successfully forwarded metrics to %s", s.Config.DestinationURL)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) forwardMetrics(data []byte, originalRequest *http.Request) error {
	// Create a new request to the destination
	req, err := http.NewRequest(http.MethodPost, s.Config.DestinationURL, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("error creating forward request: %w", err)
	}

	// Copy relevant headers
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Content-Encoding", "snappy")
	req.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")

	// Copy authorization headers if present
	for _, h := range []string{"Authorization", "X-Scope-OrgID"} {
		if v := originalRequest.Header.Get(h); v != "" {
			req.Header.Set(h, v)
			log.Printf("Forwarding header: %s", h)
		}
	}

	log.Printf("Forwarding request to %s", s.Config.DestinationURL)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Log the response status
	log.Printf("Response from destination: %d %s", resp.StatusCode, resp.Status)

	// If not successful, read and log response body
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("bad response from destination: status=%d, body=%s", resp.StatusCode, string(body))
	}

	return nil
}
