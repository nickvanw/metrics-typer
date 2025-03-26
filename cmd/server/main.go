package main

import (
	"log"
	"os"
	"time"

	"github.com/planetscale/metrics-typer/internal/config"
	"github.com/planetscale/metrics-typer/internal/server"
)

func main() {
	// Configure logging with timestamps
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("Starting metrics-typer...")

	// Log all available metric type transformations
	log.Println("Loaded metric type mappings:")
	for pattern, newType := range config.MetricTypeMap {
		log.Printf(" - Pattern '%s' -> Type '%s'", pattern, newType)
	}

	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	log.Printf("Configured with: listen_addr=%s, destination_url=%s",
		cfg.ListenAddr, cfg.DestinationURL)

	// Create and start server
	startTime := time.Now()
	s := server.New(cfg)
	log.Printf("Server initialized in %v", time.Since(startTime))

	if err := s.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
		os.Exit(1)
	}
}
