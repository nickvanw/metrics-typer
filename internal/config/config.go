package config

import (
	"fmt"
	"os"
)

// MetricTypeMap is a map of metric name patterns to their desired types
// For any metric whose name contains a key in this map, the type will be
// changed to the corresponding value
var MetricTypeMap = map[string]string{
	// PlanetScale specific mappings
	"planetscale_vttablet_queries_duration": "HISTOGRAM",
	"planetscale_vtgate_queries_duration":   "HISTOGRAM",
}

// Config holds the application configuration
type Config struct {
	ListenAddr     string
	DestinationURL string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	config := Config{
		ListenAddr:     ":8080",
		DestinationURL: "",
	}

	// Load from environment variables
	if addr := os.Getenv("LISTEN_ADDR"); addr != "" {
		config.ListenAddr = addr
	}

	if dest := os.Getenv("DESTINATION_URL"); dest != "" {
		config.DestinationURL = dest
	}

	// Validate
	if config.DestinationURL == "" {
		return nil, fmt.Errorf("destination URL is required (set DESTINATION_URL)")
	}

	return &config, nil
}
