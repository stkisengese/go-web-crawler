package main

import "testing"

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
			actual := extractDomain(tc.input)
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
			actual := calculateURLDepth(tc.input)
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
			actual := determinePageType(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
