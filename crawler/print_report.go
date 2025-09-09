package crawler

import (
	"fmt"
	"sort"
)

// Page represents a page with its normalized URL and link count
type Page struct {
	URL   string
	Count int
}

// PrintReport prints a formatted report of the crawled pages
func PrintReport(pages map[string]int, baseURL string) {
	// Print report header
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	// Sort the pages
	sortedPages := SortPages(pages)

	// Print each page with its link count
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}

// SortPages converts the map to a sorted slice
// Sorted by count (highest first), then alphabetically by URL for ties
func SortPages(pages map[string]int) []Page {
	// Convert map to slice of Page structs
	pageSlice := make([]Page, 0, len(pages))
	for url, count := range pages {
		pageSlice = append(pageSlice, Page{
			URL:   url,
			Count: count,
		})
	}

	// Sort the slice
	sort.Slice(pageSlice, func(i, j int) bool {
		// If counts are different, sort by count (highest first)
		if pageSlice[i].Count != pageSlice[j].Count {
			return pageSlice[i].Count > pageSlice[j].Count
		}
		// If counts are the same, sort alphabetically by URL
		return pageSlice[i].URL < pageSlice[j].URL
	})

	return pageSlice
}