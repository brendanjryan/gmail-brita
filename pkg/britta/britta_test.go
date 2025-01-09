package britta

import (
	"encoding/xml"
	"os"
	"strings"
	"testing"

	"github.com/brendanryan/gmail-brita/internal/config"
)

func TestFilterSetIntegration(t *testing.T) {
	tests := []struct {
		name     string
		yamlFile string
		xmlFile  string
	}{
		{
			name:     "simple filter",
			yamlFile: "../../internal/testdata/filters/simple.yaml",
			xmlFile:  "../../internal/testdata/golden/simple.xml",
		},
		{
			name:     "complex filter",
			yamlFile: "../../internal/testdata/filters/complex.yaml",
			xmlFile:  "../../internal/testdata/golden/complex.xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Load config
			cfg, err := config.LoadFromFile(tt.yamlFile)
			if err != nil {
				t.Fatalf("Failed to load config: %v", err)
			}

			// Read expected XML
			expected, err := os.ReadFile(tt.xmlFile)
			if err != nil {
				t.Fatalf("Failed to read golden file: %v", err)
			}
			expected = normalizeXML(expected)

			// Generate filters
			got, err := GenerateXML(cfg)
			if err != nil {
				t.Fatalf("GenerateXML() error = %v", err)
			}
			got = normalizeXML(got)

			// Compare
			if string(got) != string(expected) {
				t.Errorf("XML mismatch (-want +got):\n%s", diffStrings(string(expected), string(got)))
			}
		})
	}
}

// Helper functions

func normalizeXML(data []byte) []byte {
	var v interface{}
	if err := xml.Unmarshal(data, &v); err != nil {
		return data
	}
	normalized, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return data
	}
	return normalized
}

func diffStrings(expected, actual string) string {
	// Simple diff implementation
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(actual, "\n")
	var diff strings.Builder

	for i := 0; i < len(expectedLines) || i < len(actualLines); i++ {
		if i >= len(expectedLines) {
			diff.WriteString("+")
			diff.WriteString(actualLines[i])
			diff.WriteString("\n")
		} else if i >= len(actualLines) {
			diff.WriteString("-")
			diff.WriteString(expectedLines[i])
			diff.WriteString("\n")
		} else if expectedLines[i] != actualLines[i] {
			diff.WriteString("-")
			diff.WriteString(expectedLines[i])
			diff.WriteString("\n+")
			diff.WriteString(actualLines[i])
			diff.WriteString("\n")
		}
	}

	return diff.String()
}
