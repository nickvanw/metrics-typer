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
	"planetscale_vttablet_queries_duration":                       "HISTOGRAM",
	"planetscale_vtgate_queries_duration":                         "HISTOGRAM",
	"planetscale_mysql_bytes_received":                            "COUNTER",
	"planetscale_mysql_bytes_sent":                                "COUNTER",
	"planetscale_mysql_innodb_data_writes":                        "COUNTER",
	"planetscale_mysql_innodb_row_lock_time_total":                "COUNTER",
	"planetscale_mysql_innodb_row_lock_waits":                     "COUNTER",
	"planetscale_mysql_innodb_row_operations":                     "COUNTER",
	"planetscale_mysql_slow_queries":                              "COUNTER",
	"planetscale_vtgate_affected_rows":                            "COUNTER",
	"planetscale_vtgate_commands":                                 "COUNTER",
	"planetscale_vtgate_errors":                                   "COUNTER",
	"planetscale_vtgate_queries":                                  "COUNTER",
	"planetscale_vtgate_returned_rows":                            "COUNTER",
	"planetscale_vtorc_failed_recoveries":                         "COUNTER",
	"planetscale_vtorc_successful_recoveries":                     "COUNTER",
	"planetscale_vttablet_connection_pool_get":                    "COUNTER",
	"planetscale_vttablet_connection_pool_wait":                   "COUNTER",
	"planetscale_vttablet_connection_pool_wait_time_total":        "COUNTER",
	"planetscale_vttablet_errors":                                 "COUNTER",
	"planetscale_vttablet_found_rows_pool_get":                    "COUNTER",
	"planetscale_vttablet_found_rows_pool_wait":                   "COUNTER",
	"planetscale_vttablet_found_rows_pool_wait_time_total":        "COUNTER",
	"planetscale_vttablet_queries_affected_rows_total":            "COUNTER",
	"planetscale_vttablet_queries":                                "COUNTER",
	"planetscale_vttablet_queries_read_rows_total":                "COUNTER",
	"planetscale_vttablet_queries_returned_rows_total":            "COUNTER",
	"planetscale_vttablet_stream_connection_pool_get":             "COUNTER",
	"planetscale_vttablet_stream_connection_pool_wait":            "COUNTER",
	"planetscale_vttablet_stream_connection_pool_wait_time_total": "COUNTER",
	"planetscale_vttablet_transaction_pool_get":                   "COUNTER",
	"planetscale_vttablet_transaction_pool_wait":                  "COUNTER",
	"planetscale_vttablet_transaction_pool_wait_time_total":       "COUNTER",
	"planetscale_pods_iops_total":                                 "COUNTER",
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
