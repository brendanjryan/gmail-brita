package filter

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testdataPath returns an absolute path to a file in the testdata directory
func testdataPath(elem ...string) string {
	// Start with the current directory and walk up until we find testdata
	dir := "."
	for i := 0; i < 3; i++ { // Try up to 3 levels up
		path := filepath.Join(dir, "internal", "testdata", filepath.Join(elem...))
		if _, err := os.Stat(path); err == nil {
			return path
		}
		dir = filepath.Join(dir, "..")
	}
	return filepath.Join("internal", "testdata", filepath.Join(elem...))
}

func TestBuilder(t *testing.T) {
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
			expected = normalizeXML(expected)

			// Build filter set
			set := NewFilterSet([]string{"me@example.com"})
			builder := NewBuilder(set)

			// Add test filters based on test case
			switch tt.name {
			case "simple filter":
				builder.Has([]string{"test:condition"}).
					Label("test-label").
					ArchiveUnlessDirected()
			case "complex filter":
				builder.Has([]string{"test:condition", "from:test@example.com"}).
					HasNot([]string{"to:me@example.com", "cc:me@example.com"}).
					Label("test-label").
					Archive().
					MarkRead().
					Star().
					NeverSpam()

				builder = NewBuilder(set)
				builder.Has([]string{"list:test@example.com"}).
					Label("test-list").
					ArchiveUnlessDirected(WithMarkRead(true))
			}

			// Generate XML
			got, err := set.ToXML()
			if err != nil {
				t.Fatalf("ToXML() error = %v", err)
			}
			got = normalizeXML(got)

			// Compare
			if string(got) != string(expected) {
				t.Errorf("XML mismatch (-want +got):\n%s", diffStrings(string(expected), string(got)))
			}
		})
	}
}

func TestBuilderMethods(t *testing.T) {
	set := NewFilterSet([]string{"me@example.com"})
	b := NewBuilder(set)

	// Test Has
	b.Has([]string{"test:condition"})
	if len(b.filter.HasWords) != 1 || b.filter.HasWords[0] != "test:condition" {
		t.Error("Has() did not set condition correctly")
	}

	// Test HasNot
	b.HasNot([]string{"test:exclude"})
	if len(b.filter.DoesNotHaveWords) != 1 || b.filter.DoesNotHaveWords[0] != "test:exclude" {
		t.Error("HasNot() did not set exclusion correctly")
	}

	// Test Label
	b.Label("test-label")
	if len(b.filter.Labels) != 1 || b.filter.Labels[0] != "test-label" {
		t.Error("Label() did not set label correctly")
	}

	// Test Archive
	b.Archive()
	if !b.filter.Archive {
		t.Error("Archive() did not set archive flag")
	}

	// Test Star
	b.Star()
	if !b.filter.Star {
		t.Error("Star() did not set star flag")
	}

	// Test MarkRead
	b.MarkRead()
	if !b.filter.MarkRead {
		t.Error("MarkRead() did not set mark_read flag")
	}

	// Test NeverSpam
	b.NeverSpam()
	if !b.filter.NeverSpam {
		t.Error("NeverSpam() did not set never_spam flag")
	}
}

func TestArchiveUnlessDirected(t *testing.T) {
	set := NewFilterSet([]string{"me@example.com", "other@example.com"})
	b := NewBuilder(set)

	// Test basic archive unless directed
	b.Has([]string{"list:test"}).ArchiveUnlessDirected()
	if len(b.chain) != 1 {
		t.Fatal("ArchiveUnlessDirected() did not create chain filter")
	}
	archiveFilter := b.chain[0]
	if !archiveFilter.Archive {
		t.Error("ArchiveUnlessDirected() did not set archive flag")
	}
	if len(archiveFilter.DoesNotHaveWords) != 4 {
		t.Errorf("ArchiveUnlessDirected() did not set all exclusion conditions, got %d want 4", len(archiveFilter.DoesNotHaveWords))
	}

	// Test with mark read option
	b = NewBuilder(set)
	b.Has([]string{"list:test"}).ArchiveUnlessDirected(WithMarkRead(true))
	archiveFilter = b.chain[0]
	if !archiveFilter.MarkRead {
		t.Error("ArchiveUnlessDirected() did not set mark_read flag")
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
