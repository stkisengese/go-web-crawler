package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/stkisengese/go-web-crawler/crawler"
)

func main() {
	// Parse command-line arguments
	args := crawler.ParseArgs()

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
	cfg := &crawler.Config{
		Pages:              make(map[string]int),
		BaseURL:            baseURL,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, args.MaxConcurrency),
		Wg:                 &sync.WaitGroup{},
		MaxPages:           args.MaxPages,
	}

	// Start the first crawl
	startTime := time.Now()
	cfg.Wg.Add(1)
	go cfg.CrawlPage(args.URL)

	// Wait for all goroutines to complete
	cfg.Wg.Wait()
	elapsed := time.Since(startTime)

	// Print results
	crawler.PrintReport(cfg.Pages, args.URL)
	fmt.Printf("Crawl completed in %s\n", elapsed.Round(time.Millisecond))
	fmt.Printf("Average: %.2f pages/second\n", float64(len(cfg.Pages))/elapsed.Seconds())

	// Save results to CSV if specified
	if args.CSVFile != "" {
		fmt.Printf("\nExporting to CSV: %s\n", args.CSVFile)
		if err := crawler.ExportToCSV(cfg.Pages, args.URL, args.CSVFile); err != nil {
			fmt.Printf("Error exporting to CSV: %v\n", err)
		} else {
			fmt.Printf("✅ Successfully exported %d pages to %s\n", len(cfg.Pages), args.CSVFile)
		}
	}

	// Save detailed analysis to CSV if specified
	if args.DetailedCSV != "" {
		fmt.Printf("\nExporting detailed analysis to CSV: %s\n", args.DetailedCSV)
		if err := crawler.ExportDetailedCSV(cfg.Pages, args.URL, args.DetailedCSV); err != nil {
			fmt.Printf("Error exporting detailed CSV: %v\n", err)
		} else {
			fmt.Printf("✅ Successfully exported detailed analysis to %s\n", args.DetailedCSV)
		}
	}
}
