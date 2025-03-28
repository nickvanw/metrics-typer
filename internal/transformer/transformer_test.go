package transformer

import (
	"testing"

	"github.com/nickvanw/metrics-typer/internal/config"
	promproto "github.com/nickvanw/metrics-typer/internal/proto"
)

func TestGetBaseMetricName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no suffix",
			input:    "test_metric",
			expected: "test_metric",
		},
		{
			name:     "bucket suffix",
			input:    "test_metric_bucket",
			expected: "test_metric",
		},
		{
			name:     "sum suffix",
			input:    "test_metric_sum",
			expected: "test_metric",
		},
		{
			name:     "count suffix",
			input:    "test_metric_count",
			expected: "test_metric",
		},
		{
			name:     "nested suffix",
			input:    "test_count_bucket",
			expected: "test_count",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getBaseMetricName(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestSetMetricType(t *testing.T) {
	tests := []struct {
		name         string
		typeString   string
		expectedType promproto.MetricMetadata_MetricType
	}{
		{
			name:         "counter",
			typeString:   "COUNTER",
			expectedType: promproto.MetricMetadata_COUNTER,
		},
		{
			name:         "gauge",
			typeString:   "GAUGE",
			expectedType: promproto.MetricMetadata_GAUGE,
		},
		{
			name:         "histogram",
			typeString:   "HISTOGRAM",
			expectedType: promproto.MetricMetadata_HISTOGRAM,
		},
		{
			name:         "gauge histogram",
			typeString:   "GAUGEHISTOGRAM",
			expectedType: promproto.MetricMetadata_GAUGEHISTOGRAM,
		},
		{
			name:         "summary",
			typeString:   "SUMMARY",
			expectedType: promproto.MetricMetadata_SUMMARY,
		},
		{
			name:         "info",
			typeString:   "INFO",
			expectedType: promproto.MetricMetadata_INFO,
		},
		{
			name:         "stateset",
			typeString:   "STATESET",
			expectedType: promproto.MetricMetadata_STATESET,
		},
		{
			name:         "unknown",
			typeString:   "INVALID",
			expectedType: promproto.MetricMetadata_UNKNOWN,
		},
		{
			name:         "lowercase",
			typeString:   "counter",
			expectedType: promproto.MetricMetadata_COUNTER,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			transformer := New()
			metadata := &promproto.MetricMetadata{}
			transformer.setMetricType(metadata, tc.typeString)

			if metadata.Type != tc.expectedType {
				t.Errorf("Expected %v, got %v", tc.expectedType, metadata.Type)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	// Save original map and restore after test
	originalMap := config.MetricTypeMap
	defer func() {
		config.MetricTypeMap = originalMap
	}()

	// Create test config map
	config.MetricTypeMap = map[string]string{
		"test_histogram": "HISTOGRAM",
		"test_counter":   "COUNTER",
		"test_gauge":     "GAUGE",
	}

	tests := []struct {
		name              string
		inputTimeseries   []*promproto.TimeSeries
		inputMetadata     []*promproto.MetricMetadata
		expectedMetadata  map[string]promproto.MetricMetadata_MetricType
		expectedMetaCount int
		description       string
	}{
		{
			name:              "no timeseries",
			inputTimeseries:   []*promproto.TimeSeries{},
			inputMetadata:     []*promproto.MetricMetadata{},
			expectedMetadata:  map[string]promproto.MetricMetadata_MetricType{},
			expectedMetaCount: 0,
			description:       "Empty request should not create any metadata",
		},
		{
			name: "unconfigured metrics",
			inputTimeseries: []*promproto.TimeSeries{
				createTimeseries("unconfigured_metric"),
			},
			inputMetadata:     []*promproto.MetricMetadata{},
			expectedMetadata:  map[string]promproto.MetricMetadata_MetricType{},
			expectedMetaCount: 0,
			description:       "Metrics not in config should not get metadata",
		},
		{
			name: "configured metric without metadata",
			inputTimeseries: []*promproto.TimeSeries{
				createTimeseries("test_counter_value"),
			},
			inputMetadata: []*promproto.MetricMetadata{},
			expectedMetadata: map[string]promproto.MetricMetadata_MetricType{
				"test_counter_value": promproto.MetricMetadata_COUNTER,
			},
			expectedMetaCount: 1,
			description:       "Configured metrics should get metadata with correct type",
		},
		{
			name: "histogram with suffixes",
			inputTimeseries: []*promproto.TimeSeries{
				createTimeseries("test_histogram_duration_bucket"),
				createTimeseries("test_histogram_duration_sum"),
				createTimeseries("test_histogram_duration_count"),
			},
			inputMetadata: []*promproto.MetricMetadata{},
			expectedMetadata: map[string]promproto.MetricMetadata_MetricType{
				"test_histogram_duration": promproto.MetricMetadata_HISTOGRAM,
			},
			expectedMetaCount: 1,
			description:       "Histogram metrics with suffixes should get one metadata with base name",
		},
		{
			name: "existing metadata for configured metric",
			inputTimeseries: []*promproto.TimeSeries{
				createTimeseries("test_gauge_value"),
			},
			inputMetadata: []*promproto.MetricMetadata{
				createMetadata("test_gauge_value", promproto.MetricMetadata_COUNTER),
			},
			expectedMetadata: map[string]promproto.MetricMetadata_MetricType{
				"test_gauge_value": promproto.MetricMetadata_GAUGE,
			},
			expectedMetaCount: 1,
			description:       "Existing metadata should be updated to correct type",
		},
		{
			name: "mixed configured and unconfigured metrics",
			inputTimeseries: []*promproto.TimeSeries{
				createTimeseries("test_counter_total"),
				createTimeseries("unconfigured_metric"),
			},
			inputMetadata: []*promproto.MetricMetadata{},
			expectedMetadata: map[string]promproto.MetricMetadata_MetricType{
				"test_counter_total": promproto.MetricMetadata_COUNTER,
			},
			expectedMetaCount: 1,
			description:       "Only configured metrics should get metadata",
		},
		{
			name: "multiple configured metrics",
			inputTimeseries: []*promproto.TimeSeries{
				createTimeseries("test_counter_value"),
				createTimeseries("test_gauge_value"),
				createTimeseries("test_histogram_duration_bucket"),
				createTimeseries("test_histogram_duration_sum"),
				createTimeseries("test_histogram_duration_count"),
			},
			inputMetadata: []*promproto.MetricMetadata{},
			expectedMetadata: map[string]promproto.MetricMetadata_MetricType{
				"test_counter_value":      promproto.MetricMetadata_COUNTER,
				"test_gauge_value":        promproto.MetricMetadata_GAUGE,
				"test_histogram_duration": promproto.MetricMetadata_HISTOGRAM,
			},
			expectedMetaCount: 3,
			description:       "All configured metrics should get appropriate metadata",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create request with test data
			req := &promproto.WriteRequest{
				Timeseries: tc.inputTimeseries,
				Metadata:   tc.inputMetadata,
			}

			// Transform the request
			transformer := New()
			transformer.Transform(req)

			// Check the metadata count
			if len(req.Metadata) != tc.expectedMetaCount {
				t.Errorf("Expected %d metadata entries, got %d", tc.expectedMetaCount, len(req.Metadata))
			}

			// Build a map of resulting metadata for easy lookup
			resultMetadata := make(map[string]promproto.MetricMetadata_MetricType)
			for _, m := range req.Metadata {
				resultMetadata[m.MetricFamilyName] = m.Type
			}

			// Check each expected metadata entry
			for metricName, expectedType := range tc.expectedMetadata {
				actualType, exists := resultMetadata[metricName]
				if !exists {
					t.Errorf("Expected metadata for %q not found", metricName)
					continue
				}
				if actualType != expectedType {
					t.Errorf("For metric %q: expected type %v, got %v",
						metricName, expectedType, actualType)
				}
			}
		})
	}
}

// Helper function to create a timeseries with a given metric name
func createTimeseries(metricName string) *promproto.TimeSeries {
	return &promproto.TimeSeries{
		Labels: []*promproto.Label{
			{
				Name:  "__name__",
				Value: metricName,
			},
		},
		Samples: []*promproto.Sample{
			{
				Value:     1.0,
				Timestamp: 1234567890,
			},
		},
	}
}

// Helper function to create metadata with a given name and type
func createMetadata(metricName string, metricType promproto.MetricMetadata_MetricType) *promproto.MetricMetadata {
	return &promproto.MetricMetadata{
		MetricFamilyName: metricName,
		Type:             metricType,
		Help:             "Test metadata",
	}
}
