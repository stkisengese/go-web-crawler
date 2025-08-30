package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
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
	rawBaseURL := os.Args[1]
	fmt.Printf("Starting crawl of: %s\n", rawBaseURL)
	fmt.Println("===================================")

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL %s: %v\n", rawBaseURL, err)
		os.Exit(1)
	}

	maxConcurrency := 5
	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	fmt.Printf("Max concurrency: %d\n", maxConcurrency)
	fmt.Println("===================================")

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
