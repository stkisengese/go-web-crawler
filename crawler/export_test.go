package crawler

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestExtractDomain(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "domain with path",
			input:    "example.com/path/to/page",
			expected: "example.com",
		},
		{
			name:     "domain only",
			input:    "example.com",
			expected: "example.com",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "complex path",
			input:    "blog.boot.dev/posts/golang-tutorial",
			expected: "blog.boot.dev",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := ExtractDomain(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestCalculateURLDepth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "homepage",
			input:    "example.com",
			expected: 0,
		},
		{
			name:     "one level deep",
			input:    "example.com/about",
			expected: 1,
		},
		{
			name:     "two levels deep",
			input:    "example.com/blog/post",
			expected: 2,
		},
		{
			name:     "three levels deep",
			input:    "example.com/category/tech/golang",
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := CalculateURLDepth(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestDeterminePageType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "homepage",
			input:    "example.com",
			expected: "Homepage",
		},
		{
			name:     "blog page",
			input:    "example.com/blog/my-post",
			expected: "Blog",
		},
		{
			name:     "about page",
			input:    "example.com/about",
			expected: "About",
		},
		{
			name:     "contact page",
			input:    "example.com/contact-us",
			expected: "Contact",
		},
		{
			name:     "product page",
			input:    "example.com/product/widget",
			expected: "Product",
		},
		{
			name:     "generic content",
			input:    "example.com/random-page",
			expected: "Content",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := DeterminePageType(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestExportToCSV(t *testing.T) {
	// Create test data
	pages := map[string]int{
		"example.com":         5,
		"example.com/about":   3,
		"example.com/contact": 1,
	}

	filename := "test_output.csv"
	baseURL := "https://example.com"

	// Clean up test file
	defer os.Remove(filename)

	// Test CSV export
	err := ExportToCSV(pages, baseURL, filename)
	if err != nil {
		t.Fatalf("Failed to export CSV: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("CSV file was not created")
	}

	// Read and verify CSV content
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV records: %v", err)
	}

	// Check header
	if len(records) < 1 {
		t.Fatalf("CSV should have at least a header row")
	}

	header := records[0]
	expectedHeader := []string{"URL", "Internal_Link_Count", "Domain", "Base_URL", "Crawl_Timestamp"}
	if len(header) != len(expectedHeader) {
		t.Errorf("Header length mismatch: expected %d, got %d", len(expectedHeader), len(header))
	}

	// Verify we have data rows (header + data)
	if len(records) != len(pages)+1 {
		t.Errorf("Expected %d total records (header + data), got %d", len(pages)+1, len(records))
	}
}