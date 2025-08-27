package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name: "remove scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		// {
		// 	name:     "lowercase capital letters",
		// 	inputURL: "https://BLOG.boot.dev/PATH",
		// 	expected: "blog.boot.dev/path",
		// },
		// {
		// 	name:     "remove scheme and capitals and trailing slash",
		// 	inputURL: "http://BLOG.boot.dev/path/",
		// 	expected: "blog.boot.dev/path",
		// },
		// {
		// 	name:          "handle invalid URL",
		// 	inputURL:      `:\\invalidURL`,
		// 	expected:      "",
		// 	errorContains: "couldn't parse URL",
		// },
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' Fail: unexpected error: %v", i, tc.name, err)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' Fail: expected '%s', got '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
}
