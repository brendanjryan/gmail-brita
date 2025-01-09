package britta

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/brendanryan/gmail-brita/internal/config"
)

// testdataPath returns an absolute path to a file in the testdata directory
func testdataPath(elem ...string) string {
	// Try module root first (CI environment)
	moduleRoot := os.Getenv("GITHUB_WORKSPACE")
	if moduleRoot == "" {
		// Fallback to relative path for local development
		moduleRoot = "."
	}

	path := filepath.Join(moduleRoot, "internal", "testdata", filepath.Join(elem...))
	if _, err := os.Stat(path); err == nil {
		return path
	}

	// Fallback to walking up directories
	dir := "."
	for i := 0; i < 3; i++ {
		path := filepath.Join(dir, "internal", "testdata", filepath.Join(elem...))
		if _, err := os.Stat(path); err == nil {
			return path
		}
		dir = filepath.Join(dir, "..")
	}
	return filepath.Join("internal", "testdata", filepath.Join(elem...))
}

// normalizeXML normalizes XML for comparison
func normalizeXML(data []byte) ([]byte, error) {
	var v interface{}
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return xml.MarshalIndent(v, "", "  ")
}

func TestFilterSetIntegration(t *testing.T) {
	tests := []struct {
		name     string
		yamlFile string
		xmlFile  string
	}{
		{
			name:     "simple filter",
			yamlFile: testdataPath("filters", "simple.yaml"),
			xmlFile:  testdataPath("golden", "simple.xml"),
		},
		{
			name:     "complex filter",
			yamlFile: testdataPath("filters", "complex.yaml"),
			xmlFile:  testdataPath("golden", "complex.xml"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read and normalize expected XML
			expected, err := os.ReadFile(tt.xmlFile)
			if err != nil {
				t.Fatalf("Failed to read golden file: %v", err)
			}
			expectedNorm, err := normalizeXML(expected)
			if err != nil {
				t.Fatalf("Failed to normalize expected XML: %v", err)
			}

			// Load and parse YAML config
			cfg, err := config.LoadFromFile(tt.yamlFile)
			if err != nil {
				t.Fatalf("Failed to load config: %v", err)
			}

			// Generate XML
			got, err := GenerateXML(cfg)
			if err != nil {
				t.Fatalf("GenerateXML() error = %v", err)
			}
			gotNorm, err := normalizeXML(got)
			if err != nil {
				t.Fatalf("Failed to normalize generated XML: %v", err)
			}

			// Compare
			if string(gotNorm) != string(expectedNorm) {
				t.Errorf("XML mismatch\nExpected:\n%s\n\nGot:\n%s", string(expectedNorm), string(gotNorm))
			}
		})
	}
}
