package main

import (
	"reflect"
	"testing"
)

func TestGetURLFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Other</span>
		</a>
	</body>
</html>
`,
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/one",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Fatalf("unrxpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("epected %v, got %v", tc.expected, actual)
			}
		})
	}
}
