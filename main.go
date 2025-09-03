package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"
)

func main() {
	// Parse command-line arguments
	args := parseArgs()

	// Parse the base URL
	baseURL, err := url.Parse(args.URL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting crawl of: %s\n", args.URL)
	fmt.Printf("Max concurrency: %d\n", args.MaxConcurrency)
	fmt.Printf("Max pages: %d\n", args.MaxPages)
	if args.CSVFile != "" {
		fmt.Printf("CSV output file: %s\n", args.CSVFile)
	}
	if args.DetailedCSV != "" {
		fmt.Printf("Detailed CSV output file: %s\n", args.DetailedCSV)
	}
	fmt.Println("===================================")

	// Initialize the config struct
	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, args.MaxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           args.MaxPages,
	}

	// Start the first crawl
	startTime := time.Now()
	cfg.wg.Add(1)
	go cfg.crawlPage(args.URL)

	// Wait for all goroutines to complete
	cfg.wg.Wait()
	elapsed := time.Since(startTime)

	// Print results
	printReport(cfg.pages, args.URL)
	fmt.Printf("Crawl completed in %s\n", elapsed.Round(time.Millisecond))
	fmt.Printf("Average: %.2f pages/second\n", float64(len(cfg.pages))/elapsed.Seconds())

	// Save results to CSV if specified
	if args.CSVFile != "" {
		fmt.Printf("\nExporting to CSV: %s\n", args.CSVFile)
		if err := exportToCSV(cfg.pages, args.URL, args.CSVFile); err != nil {
			fmt.Printf("Error exporting to CSV: %v\n", err)
		} else {
			fmt.Printf("✅ Successfully exported %d pages to %s\n", len(cfg.pages), args.CSVFile)
		}
	}

	// Save detailed analysis to CSV if specified
	if args.DetailedCSV != "" {
		fmt.Printf("\nExporting detailed analysis to CSV: %s\n", args.DetailedCSV)
		if err := exportDetailedCSV(cfg.pages, args.URL, args.DetailedCSV); err != nil {
			fmt.Printf("Error exporting detailed CSV: %v\n", err)
		} else {
			fmt.Printf("✅ Successfully exported detailed analysis to %s\n", args.DetailedCSV)
		}
	}
}
