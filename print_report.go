package main

import "fmt"

// Page represents a page with its normalized URL and link count
type Page struct {
	URL   string
	Count int
}

// printReport prints a formatted report of the crawled pages
func printReport(pages map[string]int, baseURL string) {
	// Print report header
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	// Sort the pages
	sortedPages := sortPages(pages)

	// Print each page with its link count
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}
