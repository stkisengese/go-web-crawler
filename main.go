package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./crawler <URL> <maxConcurrency> <maxPages>")
		fmt.Println("Example: ./crawler https://example.com 5 20")
		os.Exit(1)
	}

	// Get the base URL from command line arguments
	rawBaseURL := os.Args[1]
	maxConcurrencyStr := os.Args[2]
	maxPagesStr := os.Args[3]

	// Parse maxConcurrency
	maxConcurrency, err := strconv.Atoi(maxConcurrencyStr)
	if err != nil || maxConcurrency < 1 {
		fmt.Printf("Invalid maxConcurrency '%s': must be a positive integer\n", maxConcurrencyStr)
		os.Exit(1)
	}

	// Parse maxPages
	maxPages, err := strconv.Atoi(maxPagesStr)
	if err != nil || maxPages < 1 {
		fmt.Printf("Invalid maxPages '%s': must be a positive integer\n", maxPagesStr)
		os.Exit(1)
	}

	// Parse the base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting crawl of: %s\n", rawBaseURL)
	fmt.Printf("Max concurrency: %d\n", maxConcurrency)
	fmt.Printf("Max pages: %d\n", maxPages)
	fmt.Println("===================================")

	// Initialize the config struct
	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	// Start the first crawl
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)

	// Wait for all goroutines to complete
	cfg.wg.Wait()
	// Print results
	fmt.Println("===================================")
	fmt.Printf("Crawl complete! Found %d unique pages:\n", len(cfg.pages))
	fmt.Println("===================================")

	for url, count := range cfg.pages {
		fmt.Printf("%d: %s\n", count, url)
	}
}
