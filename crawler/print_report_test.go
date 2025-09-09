package crawler

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []Page
	}{
		{
			name: "sort by count descending, then alphabetically",
			input: map[string]int{
				"example.com/page1": 3,
				"example.com/page2": 1,
				"example.com/page3": 3,
				"example.com/page4": 2,
			},
			expected: []Page{
				{URL: "example.com/page1", Count: 3},
				{URL: "example.com/page3", Count: 3},
				{URL: "example.com/page4", Count: 2},
				{URL: "example.com/page2", Count: 1},
			},
		},
		{
			name: "all same count - alphabetical sort",
			input: map[string]int{
				"example.com/zebra": 1,
				"example.com/alpha": 1,
				"example.com/beta":  1,
			},
			expected: []Page{
				{URL: "example.com/alpha", Count: 1},
				{URL: "example.com/beta", Count: 1},
				{URL: "example.com/zebra", Count: 1},
			},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []Page{},
		},
		{
			name: "single page",
			input: map[string]int{
				"example.com": 5,
			},
			expected: []Page{
				{URL: "example.com", Count: 5},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := SortPages(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}