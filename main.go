package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	
	// Get the base URL from command line arguments
	baseURL := os.Args[1]
	fmt.Printf("Starting crawl of: %s\n", baseURL)
	fmt.Println("===================================")
	
	// Initialize the pages map to track crawled pages
	pages := make(map[string]int)
	
	// Start crawling
	crawlPage(baseURL, baseURL, pages)
	
	// Print results
	fmt.Println("===================================")
	fmt.Printf("Crawl complete! Found %d unique pages:\n", len(pages))
	fmt.Println("===================================")
	
	for url, count := range pages {
		fmt.Printf("%d: %s\n", count, url)
	}
}
